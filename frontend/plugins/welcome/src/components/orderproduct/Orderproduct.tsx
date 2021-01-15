import React, { useState, Component, useEffect } from 'react';
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
import SaveIcon from '@material-ui/icons/Save';
import { DefaultApi } from '../../api/apis';
import Select from '@material-ui/core/Select';
import { EntProduct } from '../../api/models/EntProduct';
import { EntCompany } from '../../api/models/EntCompany';
import { EntTypeproduct } from '../../api/models/EntTypeproduct';
import { EntManager } from '../../api/models/EntManager';
import TextField from '@material-ui/core/TextField';
import ComponanceTable from './Tableorderproduct';
import Swal from 'sweetalert2';
//cookie
import { Cookies } from './SignInOrderproduct/Cookie'

const lightColor = 'rgba(255, 255, 255, 0.7)';

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
  }),
);

const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  width: '400px',
  padding: '100px',
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
        Poommin Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}
//testupgithub

export default function MenuAppBar() {

  var ck = new Cookies()
  var cookieEmail = ck.GetCookie()
  var cookieID = ck.GetID()
  var cookieName = ck.GetName()

  const classes = useStyles();
  const profile = { givenName: 'to Software Analysis 63' };
  const api = new DefaultApi();

  const [products, setProducts] = useState<EntProduct[]>([]);
  const [companys, setCompanys] = useState<EntCompany[]>([]);
  const [typeproducts, setTypeproducts] = useState<EntTypeproduct[]>([]);
  const [managers, setManagers] = useState<EntManager[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);


  const [managerid, setManagerid] = useState(Number);
  const [typeproductid, setTypeproductid] = useState(Number);
  const [productid, setProductid] = useState(Number);
  const [companyid, setCompanyid] = useState(Number);
  const [orderstockid, setOrderstockid] = useState(Number);
  //const [datetime, setDatetime] = useState(String);

  let stock = Number(orderstockid)
  let managerID = Number(cookieID)
  let typeproductID = Number(typeproductid)
  let productID = Number(productid)
  let companyID = Number(companyid)

  console.log(managerID)
  useEffect(() => {

    const getmanagers = async () => {

      const mn = await api.listManager({ limit: 10, offset: 0 });
      setLoading(false);
      setManagers(mn);
    };
    getmanagers();

    const getTypeproducts = async () => {

      const tp = await api.listTypeproduct({ limit: 10, offset: 0 });
      setLoading(false);
      setTypeproducts(tp);
    };
    getTypeproducts();

    const getproducts = async () => {

      const pr = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(pr);
    };
    getproducts();

    const getcompanys = async () => {

      const cp = await api.listCompany({ limit: 10, offset: 0 });
      setLoading(false);
      setCompanys(cp);
    };
    getcompanys();

  }, [loading]);

  const orderproduct = {

    managerID,
    typeproductID,
    productID,
    companyID,
    stock,
    //Addedtime :datetime   + ":00+07:00"
  }
  console.log(orderproduct)
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/orderproducts';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(orderproduct),
    };

    console.log(orderproduct); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.id != null) {
          //clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',

          }); window.setTimeout(function () { location.reload() }, 3000);
        } else {
          Toast.fire({
            icon: 'error',
            title: '<h2>บันทึกข้อมูลไม่สำเร็จ</h2>',
          });
        }
      });
  }
  const manager_id_handleChange = (event: any) => {
    setManagerid(event.target.value);
  };

  const Typeproduct_id_handleChange = (event: any) => {
    setTypeproductid(event.target.value);
  };

  const Product_id_handleChange = (event: any) => {
    setProductid(event.target.value);
  }
  const Company_id_handleChange = (event: any) => {
    setCompanyid(event.target.value);
  };
  const Orderstock_id_handleChange = (event: any) => {
    setOrderstockid(event.target.value);
  };
  /*const handleDatetimeChange = (event: any) => {
    setDatetime(event.target.value as string);
  }; */


  function HomeIcon(props: any) {
    return (
      <SvgIcon {...props}>
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
      </SvgIcon>
    );
  }

  return (
    <div className={classes.root}>
      <AppBar color="primary" position="sticky" elevation={0}>
        <Toolbar>
          <Grid container spacing={1} alignItems="center">
            <Hidden smUp>
              <Grid item>

              </Grid>
            </Hidden>
            <Grid item xs />
            <Grid item>

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
                to="/signinorderproduct">
                logout
              </Button>
            </Grid>
            <Grid item>

            </Grid>
            <Grid item>
              <IconButton color="inherit" className={classes.iconButtonAvatar}>
                <Avatar src='o' alt={cookieEmail} />
              </IconButton>
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>

      <AppBar
        component="div"
        color="primary"
        position="static"
        elevation={0}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={1}>
            <Grid item xs>
              <Typography color="inherit" variant="h2" component="h2">
                ระบบสั่งซื้อสินค้าเข้ามาในคลัง
              </Typography>
            </Grid>
            <Grid item>

            </Grid>

            <Grid item>

            </Grid>

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
        <Tabs value={0} textColor="inherit">
          <Tab textColor="inherit" label="ADD Data" />
        </Tabs>

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
            <Grid item xs={12}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ชื่อผู้จัดการ
              </Typography>
            </Grid>
            <Grid item xs={2}>

              {/*<Select
               labelId="manager_id-label"
               label="manager"
               id="manager_id"
               onChange={manager_id_handleChange}
               style = {{width: 200}}
               >{cookieID}
               {managers.map((item:EntManager)=>
               <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>)}
               </Select>*/}
              <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ชื่อสินค้า
              </Typography>
            </Grid>
            <Grid item xs={2}>

              <Select
                labelId="Equipment_id-label"
                label="Equipment"
                id="Equipment_id"
                onChange={Product_id_handleChange}
                style={{ width: 200 }}
              >
                {products.map((item: EntProduct) =>
                  <MenuItem key={item.id} value={item.id}>{item.nameProduct}</MenuItem>)}
              </Select>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ประเภทของสินค้า
              </Typography>
            </Grid>
            <Grid item xs={2}>

              <Select
                labelId="medicalType_id-label"
                label="medicalType"
                id="medicalType_id"
                onChange={Typeproduct_id_handleChange}
                style={{ width: 200 }}
              >
                {typeproducts.map((item: EntTypeproduct) =>
                  <MenuItem key={item.id} value={item.id}>{item.typeproduct}</MenuItem>)}
              </Select>



            </Grid>


            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                บริษัท
              </Typography>
            </Grid>
            <Grid item xs={2}>

              <Select
                labelId="Equipment_id-label"
                label="Equipment"
                id="Equipment_id"
                onChange={Company_id_handleChange}
                style={{ width: 200 }}
              >
                {companys.map((item: EntCompany) =>
                  <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>)}
              </Select>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>


            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                จำนวน/ชิ้น
              </Typography>
            </Grid>
            <Grid item xs={2}>

              <TextField id="outlined-number" type='number' InputLabelProps={{
                shrink: true,
              }} label="กรุณาใส่จำนวน" variant="outlined"
                onChange={Orderstock_id_handleChange}
              />


            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            {/* <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                วันที่/เวลาที่บันทึก
              </Typography>
            </Grid>
            <Grid item xs={2}>
               <form  noValidate>
               <TextField
                      id="date"
                      label="DateTime"
                      type="datetime-local"
                      value={datetime}
                      onChange={handleDatetimeChange}
                      //defaultValue="2017-05-24"
                      className={classes.textField}
                      InputLabelProps={{
                        shrink: true,
                      }}
                    />
              </form>
              
                  
            </Grid>
            <Grid item xs={2}></Grid>
                    <Grid item xs={2}> </Grid>*/}

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
                  save();
                }}

                startIcon={<SaveIcon
                />}
              >
                Save
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
      <ComponanceTable></ComponanceTable>

    </div>
  );
}