import React, { FC, useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
//import ComponanceTable from '../EmployeeWorkingHours';
import { Link as RouterLink } from 'react-router-dom';
import moment from "moment";

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  Avatar,
  Button,
  TextField,
  Typography,
  Link,
  SvgIcon,
  AppBar,
  Toolbar,
  IconButton,
  Hidden,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import { EntDay } from '../../api/models/EntDay'; // import interface Day
import { EntRole } from '../../api/models/EntRole'; // import interface Role
import { EntBeginWork } from '../../api/models/EntBeginWork'; // import interface EntBeginWork
import { EntGetOffWork } from '../../api/models/EntGetOffWork'; // import interface EntBeginWork

//import { Cookies } from './orderproduct/SignInOrderproduct/Cookie';
import SearchIcon from '@material-ui/icons/Search';


const lightColor = 'rgba(255, 255, 255, 0.7)';
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
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
  button: {
    borderColor: lightColor,
  },
  iconButtonAvatar: {
    padding: 4,
  },
  secondaryBar: {
    zIndex: 0,
  },

}));

interface EmployeeWorkingHours {
  CodeWork:   string;
  IDNumber:   string;
  Employee:   number;
  Day:        number;
  Role:       number;
  BeginWork:  number;
  GetOffWork: number;
  Wages:      number;
  // create_by: number;
}

const EmployeeWorkingHours: FC<{}> = () => {

  //var ck = new Cookies()
  //var cookieEmail = ck.GetCookie()
  //var cookieID = ck.GetID()
  //var cookieName = ck.GetName()

  const classes = useStyles();
  const http = new DefaultApi();

  const [EmployeeWorkingHourss, setEmployeeWorkingHourss] = React.useState<
    Partial<EmployeeWorkingHours>
  >({});

  const [loading, setLoading] = useState(true);
  const [Days, setDays] = React.useState<EntDay[]>([]);
  const [Roles, setRoles] = React.useState<EntRole[]>([]);
  const [Employees, setEmployees] = React.useState<EntEmployee[]>([]);
  const [BeginWorks, setBeginWorks] = React.useState<EntBeginWork[]>([]);
  const [GetOffWorks, setGetOffWorks] = React.useState<EntGetOffWork[]>([]);

  //const [Day, SetDayid] = useState(Number);
  //const [Role, SetRolesid] = useState(Number);
  //const [Employee, SetEmployeeid] = useState(Number);
  //const [Shift, SetShiftid] = useState(Number);

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'center-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  function Copyright() {
    return (
      <Typography variant="body2" color="inherit" align="center">
        {'Copyright © '}
        <Link color="inherit" href="https://material-ui.com/">
          Thanabodee Website
        </Link>{' '}
        {new Date().getFullYear()}
        {'.'}
      </Typography>
    );
  }

  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setLoading(false);
    setEmployees(res);
  };

  const getDay = async () => {
    const res = await http.listDay({ limit: 10, offset: 0 });
    setLoading(false);
    setDays(res);
  };

  const getRole = async () => {
    const res = await http.listRole({ limit: 10, offset: 0 });
    setLoading(false);
    setRoles(res);
  };

  const getBeginWork = async () => {
    const res = await http.listBeginwork({ limit: 10, offset: 0 });
    setLoading(false);
    setBeginWorks(res);
  };

  const getGetOffWork = async () => {
    const res = await http.listGetoffwork({ limit: 10, offset: 0 });
    setLoading(false);
    setGetOffWorks(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getEmployee();
    getDay();
    getRole();
    getBeginWork();
    getGetOffWork();
  }, [loading]);

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

  /*const IDEmployeehandleChange = (event: any) => {
    SetIDEmployee(event.target.value as string);
  };
  const IDNumberhandleChange = (event: any) => {
    SetIDNumber(event.target.value as string);
  };
  const WagesThandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    SetWages(event.target.value as number);
  };*/

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof EmployeeWorkingHours;
    const { value } = event.target;
    setEmployeeWorkingHourss({ ...EmployeeWorkingHourss, [name]: value });
    console.log(EmployeeWorkingHourss);
  };

  const alertMessage = (icon: any, title: any) => {
      Toast.fire({
        icon: icon,
        title: title,
      });
    }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'IDEmployee':
        alertMessage("error","รหัสพนักงานขึ้นต้นด้วย A,B,C ตามด้วยเลข 5 หลัก");
        return;
      case 'IDNumber':
        alertMessage("error","กรุณากรอกเลขบัตรประชาชน 13 หลักให้ถูกต้อง");
        return;
      case 'Wages':
        alertMessage("error","กรุณาระบุค่าจ้างให้เป็นตัวเลขและไม่ติดลบ");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  // clear input form
  function clear() {
    setEmployeeWorkingHourss({});
  }

  // function save data

  console.log(EmployeeWorkingHourss)
  const save = async () => {
      if(EmployeeWorkingHourss.Wages){
          var Wages : number = +EmployeeWorkingHourss.Wages;
          EmployeeWorkingHourss.Wages = Wages;        
      }
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
          checkCaseSaveError(data.error.Name)
        }
      });
    }

    function HomeIcon(props: any) {
      return (
        <SvgIcon {...props}>
          <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
        </SvgIcon>
      );
    }
    
  return (
    <div className={classes.root}>
      <AppBar component="div" color="primary" position="sticky" elevation={0}>
        <Toolbar>
          <Grid container spacing={1} alignItems="center">
            <Grid item>
                <Typography color="inherit" variant="h3" component="h3">
                  ระบบตารางเวลาทำงานของพนักงาน
                </Typography>
            </Grid>
            <Hidden smUp>
              <Grid item>

              </Grid>
            </Hidden>
            <Grid item xs />

              <Grid item>
                <IconButton
                  style={{ marginLeft: 20 }}
                  component={RouterLink}
                  to="/"
                >
                  <HomeIcon color="inherit" />
                </IconButton>
              </Grid>

              <Grid item>
                <Button className={classes.button} variant="outlined" color="inherit"
                  size="small" component={RouterLink}
                  to="/signinorderproduct">
                  logout
                </Button>
              </Grid>
            
              <Grid item>
                <IconButton color="inherit" className={classes.iconButtonAvatar}>
                  <Avatar src='o' />
              </IconButton>
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>

      <AppBar
        component="div"
        color="inherit"
        className={classes.secondaryBar}
        position="static"
        elevation={1}
      >
        <Toolbar>
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

            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เลขบัตรประชาชน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                  id="IDNumber" 
                  name = "IDNumber"
                  label = "ใส่เลขบัตรประชาชน"
                  InputLabelProps={{shrink: true,}} 
                  variant="outlined"
                  onChange={handleChange}
                  value={EmployeeWorkingHourss.IDNumber || ''}   
              />
              </FormControl>
            </Grid>

            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>รหัสพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                  id="CodeWork"
                  name ="CodeWork"
                  label = "ใส่โค๊ดทำงาน"
                  InputLabelProps={{shrink: true,}} 
                  variant="outlined"
                  onChange={handleChange}
                  value={EmployeeWorkingHourss.CodeWork || ''}
                  
              />
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
              <div className={classes.paper}>เวลาเริ่มงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกเวลา เริ่มงาน</InputLabel>
                <Select
                  name="BeginWork"
                  value={EmployeeWorkingHourss.BeginWork || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {BeginWorks.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {moment(item.beginWork).format("LT")}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เวลาเลิกงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกเวลา เลิกงาน</InputLabel>
                <Select
                  name="GetOffWork"
                  value={EmployeeWorkingHourss.GetOffWork || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {GetOffWorks.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {moment(item.getOffWork).format("LT")}
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

            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ค่าจ้าง</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                  id="Wages" 
                  name = "Wages"
                  label = "ใส่ค่าจ้าง"
                  InputLabelProps={{shrink: true,}} 
                  variant="outlined"
                  onChange={handleChange}
                  value={EmployeeWorkingHourss.Wages || ''}
              />
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
      </Toolbar>
      </AppBar>
    </div>
  );
};

export default EmployeeWorkingHours;
