import React, { FC, useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
//import ComponanceTable from '../EmployeeWorkingHours';
import { Link as RouterLink } from 'react-router-dom';

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  Avatar,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import { EntDay } from '../../api/models/EntDay'; // import interface Day
import { EntRole } from '../../api/models/EntRole'; // import interface Role
import { EntShift } from '../../api/models/EntShift'; // import interface Shift

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

interface EmployeeWorkingHours {
  Employee: number;
  Day: number;
  Role: number;
  Shift: number;
  // create_by: number;
}

const EmployeeWorkingHours: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [EmployeeWorkingHourss, setEmployeeWorkingHourss] = React.useState<
    Partial<EmployeeWorkingHours>
  >({});

  const [Days, setDays] = React.useState<EntDay[]>([]);
  const [Roles, setRoles] = React.useState<EntRole[]>([]);
  const [Employees, setEmployees] = React.useState<EntEmployee[]>([]);
  const [Shifts, setShifts] = React.useState<EntShift[]>([]);

  //const [Day, SetDayid] = useState(Number);
  //const [Role, SetRolesid] = useState(Number);
  //const [Employee, SetEmployeeid] = useState(Number);
  //const [Shift, SetShiftid] = useState(Number);

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setEmployees(res);
  };

  const getDay = async () => {
    const res = await http.listDay({ limit: 10, offset: 0 });
    setDays(res);
  };

  const getRole = async () => {
    const res = await http.listRole({ limit: 10, offset: 0 });
    setRoles(res);
  };

  const getShift = async () => {
    const res = await http.listShift({ limit: 10, offset: 0 });
    setShifts(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getEmployee();
    getDay();
    getRole();
    getShift();
  }, []);

  //const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  //  SetEmployeeid(event.target.value as number);
  //};
  //const DayhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  //  SetDayid(event.target.value as number);
  //};
  //const RolehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  //  SetRolesid(event.target.value as number);
  //};
  //const ShifthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  //  SetShiftid(event.target.value as number);
  //};

  // set data to object playlist_video
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof EmployeeWorkingHours;
    const { value } = event.target;
    setEmployeeWorkingHourss({ ...EmployeeWorkingHourss, [name]: value });
    console.log(EmployeeWorkingHourss);
  };


  // clear input form
  function clear() {
    setEmployeeWorkingHourss({});
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/employeeworkinghourss';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(EmployeeWorkingHourss),
    };

    console.log(EmployeeWorkingHourss); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`Employee Working Hours`}>
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10 }}>Thanabodee Petchrey</div>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>พนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกพนักงาน</InputLabel>
                <Select
                  name="Employee"
                  value={EmployeeWorkingHourss.Employee || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {Employees.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่ทำงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกวันที่ทำงาน</InputLabel>
                <Select
                  name="Day"
                  value={EmployeeWorkingHourss.Day || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {Days.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.day}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เวลา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกเวลา</InputLabel>
                <Select
                  name="Shift"
                  value={EmployeeWorkingHourss.Shift || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {Shifts.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>หน้าที่ที่รับผิดชอบ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกหน้าที่ที่รับผิดชอบ</InputLabel>
                <Select
                  value={EmployeeWorkingHourss.Role || ''} // (undefined || '') = ''
                  onChange={handleChange}
                  name="Role"
                >
                  {Roles.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.role}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={3}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                Save
              </Button>
            </Grid>
            <Grid item xs={3}>
              <Button
                component={RouterLink}
                to="/SplitsystemManager"
                variant="contained"
                color="secondary"
                size="large"
              >
              Back
             </Button>
            </Grid>
            <Grid item xs={2}>
              <Button
                component={RouterLink}
                to="/TableEmployeeWorkingHours"
                variant="contained"
                color="secondary"
                size="large"
              >
                Table
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default EmployeeWorkingHours;