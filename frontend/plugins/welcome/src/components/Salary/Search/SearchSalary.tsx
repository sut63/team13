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
import { FormControl, InputLabel, TextField } from '@material-ui/core';
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
  const [employeeid, setEmployeeid] = useState(String);
  const [search, setSearch] = useState(false);
  const [checkSalaryName, setSalaryNames] = useState(false);

  /*let employeeID = Number(cookieID)*/
  //let employeeID = String(employeeid)
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }



  const [salarys, setSalarys] = useState<EntSalary[]>([]);


  const deleteSystemequipments = async (id: number) => {
    const res = await api.deleteSalary({ id: id });
    setLoading(true);
  };
  useEffect(() => {
    const getEmployees = async () => {
      const em = await api.listEmployee({ offset: 0 });
      setLoading(false);
      setEmployees(em);
    };
    getEmployees();

    const getSalary = async () => {
      const sa = await api.listSalary({ offset: 0 });
      setLoading(false);
      setSalarys(sa);
    };
    getSalary();
  }, [loading]);

  const Employee_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployeeid(event.target.value as string);
    setSalaryNames(false);
    setSearch(false);
  }
  
  const getCheckSalary = async () => {
    var check = false;
    salarys.map(item => {
      if (employeeid != "") {
        if (item.edges?.employee?.name?.startsWith(employeeid)) {
          setSalaryNames(true);
          alertMessage("success", "ค้นหาข้อมูลสำเร็จ");
          check = true;
        }
      }     
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูลที่ค้นหา");
    }
    console.log(checkSalaryName)
    if (employeeid == "") {
      alertMessage("info", "กรอกชื่อพนักงาน");
      window.setTimeout(function(){location.reload()},1000);
    }
  };

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

            <Grid item xs={3}></Grid>
            
            <Grid item xs={2}>
              <Typography color="secondary" variant="h6" component="h1">
                ชื่อพนักงาน
              </Typography>
              
            </Grid>
            <Grid item xs={2}>
            <FormControl
                fullWidth
                variant="outlined"

              ><TextField
                  id="employee"
                  label="กรอกชื่อพนักงาน"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={employeeid}
                  onChange={Employee_id_handleChange}
                  style={{ marginRight: 300, width: 300 }}
                />
              </FormControl>
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
                  getCheckSalary();
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

      <Grid item xs={12}>
        <Paper>
          {search ? (
            <div>
              {  checkSalaryName ? (
                <TableContainer component={Paper}>
                  <Table className={classes.table} aria-label="simple table">
                    <TableHead>
                      <TableRow>
                        <TableCell align="center">หมายเลข</TableCell>
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
                      {salarys.filter((filter: any) =>
                        filter.edges?.employee?.name.startsWith(employeeid)).map((item: any) => (
                          <TableRow key={item.id}>
                            <TableCell align="center">{item.id}</TableCell>
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
              )
                : employeeid == "" ? (
                  <div>
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                            <TableCell align="center">หมายเลข</TableCell>
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
                          {salarys.filter((filter: any) =>
                            filter.edges?.employee?.nameProduct.startsWith(employeeid)).map((item: any) => (
                              <TableRow key={item.id}>
                                <TableCell align="center">{item.id}</TableCell>
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
                ) : null}
            </div>
          ) : null}
        </Paper>
      </Grid>

    </div>
  );
}