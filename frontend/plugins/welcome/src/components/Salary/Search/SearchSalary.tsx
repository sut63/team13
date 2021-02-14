import React, { useState, useEffect } from 'react';
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
import SvgIcon from '@material-ui/core/SvgIcon';
import { DefaultApi } from '../../../api/apis';
import Select from '@material-ui/core/Select';
import { EntEmployee } from '../../../api/models/EntEmployee';
import Swal from 'sweetalert2';
import SearchIcon from '@material-ui/icons/Search';
import { EntSalary } from '../../../api/models/EntSalary';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
import { Cookies } from '../../orderproduct/SignInOrderproduct/Cookie'
import AccountCircle from '@material-ui/icons/AccountCircle';

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
    table: {
      minWidth: 650,
    },
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
  }),
);


/*function Copyright() {
  return (
    <Typography variant="body2" color="inherit" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Poommin Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}*/
//testupgithub

export default function MenuAppBar() {

  /*var ck = new Cookies()
  var cookieEmail = ck.GetCookie()
  var cookieID = ck.GetID()
  var cookieName = ck.GetName()*/
  const classes = useStyles();
  const api = new DefaultApi();

  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [loading, setLoading] = useState(true);
  const [employeeid, setEmployeeid] = useState(Number);

  /*let employeeID = Number(cookieID)*/
  let employeeID = Number(employeeid)



  const [salarys, setSalarys] = useState<EntSalary[]>();


  const deleteSystemequipments = async (id: number) => {
    const res = await api.deleteSalary({ id: id });
    setLoading(true);
  };
  useEffect(() => {
    const getEmployees = async () => {
      const em = await api.listEmployee({ limit: 10, offset: 0 });
      setLoading(false);
      setEmployees(em);
    };
    getEmployees();
  }, [loading]);

  const salary = {
    employeeID,
  }
  console.log(salary)

  const Employee_id_handleChange = (event: any) => {
    setEmployeeid(event.target.value);
  }
  var lenSalary: number
  
  const getCheckinsorder = async () => {
    const res = await api.getSalary({ id: employeeid })
    setSalarys(res)
    lenSalary = res.length
    if (lenSalary > 0) {
      //setOpen(true)
      Toast.fire({
        icon: 'success',
        title: 'ค้นหาข้อมูลสำเร็จ',
      })
    } else {
      //setFail(true)
      Toast.fire({
        icon: 'error',
        title: 'ค้นหาข้อมูลไม่พบ',
      })
    }
  }

  function HomeIcon(props: any) {
    return (
      <SvgIcon {...props}>
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
      </SvgIcon>
    );
  }

  var ck = new Cookies()
  //var cookieEmail = ck.GetCookie()
  //var cookieID = ck.GetID()
  var cookieName = ck.GetName()

  return (
    <div className={classes.root}>
      <AppBar
        component="div"
        color= 'secondary'
        position="static"
        elevation={0}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={1}>
            <Grid item xs>
              <Typography color="inherit" variant="h3" component="h2">
                ระบบค้นหาบันทึกเงินเดือนพนักงาน
              </Typography>
            </Grid>
            <Grid item>
            <IconButton
                aria-label="account of current user"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                color="inherit"
              >
                <AccountCircle />
              </IconButton>
                </Grid>
            <Grid item>
              
            <Typography color="inherit" variant="h6" component="h2">
                {cookieName}
              </Typography>
                </Grid>
            <Grid item>
                <IconButton 
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/SplitsystemManager"
                >    
                <HomeIcon color="inherit" />
                </IconButton>
                </Grid>
                <Grid item>
                  
            <Button className={classes.button} variant="outlined" color="inherit" 
            size="small" component={RouterLink}
            to="/">
                logout
              </Button>
                </Grid>  
            
            
          </Grid>
        </Toolbar>
      </AppBar>


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
            <Grid item xs={12}>

            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>

            
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="secondary" variant="h6" component="h1">
                ชื่อพนักงาน
              </Typography>
            </Grid>
            <Grid item xs={2}>

              <Select
                labelId="Equipment_id-label"
                label="Equipment"
                id="Equipment_id"
                onChange={Employee_id_handleChange}
                style={{ width: 200 }}
              >
                {employees.map((item: EntEmployee) =>
                  <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>)}
              </Select>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>


            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Button

                variant="contained"
                color="secondary"
                size="large"
                className={classes.button}
                onClick={() => {
                  getCheckinsorder();
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

      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell align="center">หมายเลข</TableCell>
              <TableCell align="center">รหัสพนักงาน</TableCell>
              <TableCell align="center">เลขบัญชีธนาคาร</TableCell>
              <TableCell align="center">รายชื่อพนักงาน</TableCell>
              <TableCell align="center">ตำแหน่ง</TableCell>
              <TableCell align="center">ผลการประเมิน</TableCell>
              <TableCell align="center">เงินเดือน</TableCell>
              <TableCell align="center">โบนัส</TableCell>
              <TableCell align="center">วันที่เงินเดือนออก</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {salarys === undefined
              ? null
              : salarys.map((item: any) => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.iDEmployee}</TableCell>
                  <TableCell align="center">{item.accountNumber}</TableCell>
                  <TableCell align="center">{item.edges?.employee?.name}</TableCell>
                  <TableCell align="center">{item.edges?.position?.position}</TableCell>
                  <TableCell align="center">{item.edges?.assessment?.assessment}</TableCell>
                  <TableCell align="center">{item.bonus}</TableCell>
                  <TableCell align="center">{item.salary}</TableCell>
                  <TableCell align="center">{moment(item.salaryDatetime).format('DD/MM/YYYY HH:mm:ss')}</TableCell>

                  <TableCell align="center">
                    <Button
                      onClick={() => {
                        deleteSystemequipments(item.id);
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

    </div>
  );
}
