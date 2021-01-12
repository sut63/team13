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
import { EntSalary } from '../../api/models/EntSalary';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [salarys, setSalarys] = React.useState<EntSalary[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getSalary = async () => {
     const res = await api.listSalary({ limit: 10, offset: 0 });
     setLoading(false);
     setSalarys(res);
   };
   getSalary();
 }, [loading]);
 
 const deleteSalary = async (id: number) => {
   const res = await api.deleteSalary({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">หมายเลข</TableCell>
           <TableCell align="center">รายชื่อพนักงาน</TableCell>
           <TableCell align="center">ตำแหน่ง</TableCell>
           <TableCell align="center">ผลการประเมิน</TableCell>
           <TableCell align="center">เงินเดือน</TableCell>
           <TableCell align="center">โบนัส</TableCell>
           <TableCell align="center">วันที่เงินเดือนออก</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {salarys.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.employee?.name}</TableCell>
             <TableCell align="center">{item.edges?.position?.position}</TableCell>
             <TableCell align="center">{item.edges?.assessment?.assessment}</TableCell>
             <TableCell align="center">{item.bonus}</TableCell>
             <TableCell align="center">{item.salary}</TableCell>
             <TableCell align="center">{item.salaryDatetime}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                  deleteSalary(item.id);
                 }}
                 style={{ marginLeft: 20 }}
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
