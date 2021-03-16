import React, { useState,Component, useEffect } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import Grid from '@material-ui/core/Grid';
import MenuItem from '@material-ui/core/MenuItem';
import Button from '@material-ui/core/Button';
import Hidden from '@material-ui/core/Hidden';
import { Link as RouterLink } from 'react-router-dom';
import Link from '@material-ui/core/Link';
import Avatar from '@material-ui/core/Avatar';
import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';
import { DefaultApi } from '../../../api/apis';
import Select from '@material-ui/core/Select';
import { EntEmployee } from '../../../api/models/EntEmployee';
import { EntEmployeeWorkingHours } from '../../../api/models/EntEmployeeWorkingHours';
import Swal from 'sweetalert2';
import SearchIcon from '@material-ui/icons/Search';
import moment from 'moment';
import SvgIcon from '@material-ui/core/SvgIcon'; // Api Gennerate From Command
import Paper from '@material-ui/core/Paper';
//import { Cookies } from '../../../SignInOrderproduct/Cookie'
//import { ContentHeader } from '@backstage/core';
import { createMuiTheme } from '@material-ui/core/styles';
import purple from '@material-ui/core/colors/purple';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TextField from '@material-ui/core/TextField';

const lightColor = 'rgba(255, 255, 255, 0.7)';

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

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    iconButtonAvatar: {
      padding: 4,
    },
    secondaryBar: {
      zIndex: 0,
    },
    button: {
      borderColor: lightColor,
    },
    title: {
      flexGrow: 1,
    },
    textField: {
      width: 200,
    },
    table: {
      minWidth: 650,
    },
    primary: {
      light: '#757ce8',
      main: '#3f50b5',
      dark: '#002884',
      contrastText: '#fff',
    },
  }),
);

const theme = createMuiTheme({
  palette: {
    primary: {
      main: purple[500],
    },
    secondary: {
      main: '#f44336',
    },
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

export default function MenuAppBar() {

  //var ck = new Cookies()
  //var cookieEmail = ck.GetCookie()
  //var cookieID = ck.GetID()
  //var cookieName = ck.GetName()
  const classes = useStyles();
  const api = new DefaultApi();
  const [search, setSearch] = useState(false);
  const [checkEmployeeName, setEmployeeName] = useState(false);
  const [Employees, setEmployees] = useState<EntEmployee[]>([]);
  const [loading, setLoading] = useState(true);
  const [Employeeid, setEmployeeid] = useState(String);
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }


  const [EmployeeWorkingHourss, setEmployeeWorkingHourss] = useState<EntEmployeeWorkingHours[]>([]);


  useEffect(() => {
    const getEmployees = async () => {
      const em = await api.listEmployee({ limit: 10, offset: 0 });
      setLoading(false);
      setEmployees(em);
    };
    getEmployees();
    const getEmployeeWorkingHours = async () => {
      const ewh = await api.listEmployeeworkinghours({ limit: 10, offset: 0 });
      setLoading(false);
      setEmployeeWorkingHourss(ewh);
    };
    getEmployeeWorkingHours();
  }, [loading]);

  const Employee_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployeeid(event.target.value as string);
    setEmployeeName(false);
    setSearch(false);
  }
  
  const getCheckEmployeeWorkingHours = async () => {
    var check = false;
    EmployeeWorkingHourss.map(item => {
      if (Employeeid != "") {
        if (item.edges?.employee?.name?.startsWith(Employeeid)) {
          setEmployeeName(true);
          alertMessage("success", "ค้นหาข้อมูลสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูลที่ค้นหา");
    }
    console.log(checkEmployeeName)
    if (Employeeid == "") {
      alertMessage("info", "กรอกชื่อพนักงาน");
    }
  };

 function HomeIcon(props:any) {
    return (
      <SvgIcon {...props}>
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
      </SvgIcon>
    );
  }

  return (
    <div className={classes.root}>
      <AppBar component="div" color= 'primary' position="static" elevation={0}>
        <Toolbar>
          <Grid container alignItems="center" spacing={1}>
            <Hidden smUp>
              <Grid item>

              </Grid>
            </Hidden>
            <Grid item xs>
              <Typography color="inherit" variant="h3" component="h2">
                ระบบตารางเวลาทำงานของพนักงาน
              </Typography>
            </Grid>
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
              to="/SignInOrderproduct">
                logout
              </Button>
            </Grid>  
            
            <Grid item>
              <Button className={classes.button} variant="outlined" color="inherit"
                size="small" component={RouterLink}
                to="/EmployeeWorkingHours">
                Back
              </Button>
            </Grid>
          </Grid>
        </Toolbar>
        
      <AppBar
        component="div"
        className={classes.secondaryBar}
        color="inherit"
        position="static"
        elevation={1}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={4}>
            <Grid item xs={12}></Grid>
            <Grid item xs={12}></Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ชื่อพนักงานที่ต้องการค้นหา
              </Typography>
            </Grid>
            <Grid item xs={2}>

              <TextField
                  id="Employee_id"
                  label="Employee"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={Employeeid}
                  onChange={Employee_id_handleChange}
                  style={{ marginRight: 300, width: 300 }}
                />
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>



            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Button

                variant="contained"
                color="primary"
                size="large"
                className={classes.button}
                onClick={() => {
                  getCheckEmployeeWorkingHours();
                  setSearch(true);
                }}

                startIcon={<SearchIcon
                />}
              >
                Search
              </Button>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>
            <Grid item xs={12}></Grid>
            <Grid item xs={12}></Grid>

          </Grid>
        </Toolbar>
      </AppBar>
      <AppBar
        component="div"
        className={classes.secondaryBar}
        color="primary"
        position="static"
        elevation={0}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={1}>
            <Grid item xs>
              <Copyright />

            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>
      {search ? (
        <div>
          {  checkEmployeeName ? (
      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="EmployeeWorkingHours Table">
          <TableHead>
            <TableRow>
              <TableCell align="center">No.</TableCell>
              <TableCell align="center">ชื่อพนักงาน</TableCell>
              <TableCell align="center">เลขบัตรประชาชน</TableCell>
              <TableCell align="center">รหัสพนักงาน</TableCell>
              <TableCell align="center">วันที่เข้าทำงาน</TableCell>
              <TableCell align="center">เวลาเริ่มงาน</TableCell>
              <TableCell align="center">เวลาเลิกงาน</TableCell>
              <TableCell align="center">หน้าที่ที่รับผิดชอบ</TableCell>
              <TableCell align="center">ค่าจ้าง</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {EmployeeWorkingHourss.filter((filter: any) =>
              filter.edges?.employee?.name.startsWith(Employeeid)).map((item: any) => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.employee?.name}</TableCell>
                  <TableCell align="center">{item.iDNumber}</TableCell>
                  <TableCell align="center">{item.codeWork}</TableCell>
                  <TableCell align="center">{item.edges?.day?.day}</TableCell>
                  <TableCell align="center">{moment(item.edges?.startWork?.startWork).format("LT")}</TableCell>
                  <TableCell align="center">{moment(item.edges?.endWork?.endWork).format("LT")}</TableCell>
                  <TableCell align="center">{item.edges?.role?.role}</TableCell>
                  <TableCell align="center">{item.wages}</TableCell>
                </TableRow>
              ))}
          </TableBody>
        </Table>
      </TableContainer>
         ) : null}
         </div>
       ) : null}
        
      </AppBar>
      
            
    </div>
  );
 }