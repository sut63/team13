import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntEmployeeWorkingHours } from '../../api/models/EntEmployeeWorkingHours';
import moment from "moment";
const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
});

export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [EmployeeWorkingHourss, setEmployeeWorkingHourss] = useState<EntEmployeeWorkingHours[]>(Array);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getUsers = async () => {
      const res = await api.listEmployeeworkinghours({ limit: 10, offset: 0 });
      setLoading(false);
      setEmployeeWorkingHourss(res);
    };
    getUsers();
  }, [loading]);

  const deleteEmployeeWorkingHours = async (id: number) => {
    const res = await api.deleteEmployeeworkinghours({ id: id });
    setLoading(true);
  };

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">No.</TableCell>
            <TableCell align="center">รหัสพนักงาน</TableCell>
            <TableCell align="center">เลขบัตรประชาชน</TableCell>
            <TableCell align="center">ชื่อพนักงาน</TableCell>
            <TableCell align="center">วันที่เข้าทำงาน</TableCell>
            <TableCell align="center">เวลาเริ่มทำงาน</TableCell>
            <TableCell align="center">เวลาเลิกงาน</TableCell>
            <TableCell align="center">หน้าที่ที่รับผิดชอบ</TableCell>
            <TableCell align="center">ค่าจ้าง</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {EmployeeWorkingHourss === undefined
            ? null
            : EmployeeWorkingHourss.map((item:any) => (
              <TableRow key={item.id}>
                <TableCell align="center">{item.id}</TableCell>
                <TableCell align="center">{item.iDEmployee}</TableCell>
                <TableCell align="center">{item.iDNumber}</TableCell>
                <TableCell align="center">{item.edges?.employee?.name}</TableCell>
                <TableCell align="center">{item.edges?.day?.day}</TableCell>
                <TableCell align="center">{moment(item.edges?.shift?.timeStart).format("LT")}</TableCell>
                <TableCell align="center">{moment(item.edges?.shift?.timeEnd).format("LT")}</TableCell>
                <TableCell align="center">{item.edges?.role?.role}</TableCell>
                <TableCell align="center">{item.wages}</TableCell>
                <TableCell align="center">
                  <Button
                    onClick={() => {
                      deleteEmployeeWorkingHours(item.id);
                    }}
                    style={{ marginLeft: 10 }}
                    variant="contained"
                    color="secondary"
                  >
                    Delete
               </Button>
                </TableCell>
              </TableRow>
            ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
