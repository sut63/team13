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
import { EntProduct } from '../../../api/models/EntProduct';
import { EntCompany } from '../../../api/models/EntCompany';
import { EntTypeproduct } from '../../../api/models/EntTypeproduct';
import TextField from '@material-ui/core/TextField';
import ComponanceTable from '../Tableorderproduct';
import Swal from 'sweetalert2';
//cookie
import { Cookies } from '../SignInOrderproduct/Cookie'
import SearchIcon from '@material-ui/icons/Search';



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
  //const profile = { givenName: 'to Software Analysis 63' };
  const api = new DefaultApi();

  const [products, setProducts] = useState<EntProduct[]>([]);
  const [companys, setCompanys] = useState<EntCompany[]>([]);
  const [typeproducts, setTypeproducts] = useState<EntTypeproduct[]>([]);
  //const [managers, setManagers] = useState<EntManager[]>([]);
  //const [status, setStatus] = useState(false);
  //const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);


  //const [managerid, setManagerid] = useState(Number);
  const [typeproductid, setTypeproductid] = useState(Number);
  const [productid, setProductid] = useState(Number);
  const [companyid, setCompanyid] = useState(Number);
  const [orderstockid, setOrderstockid] = useState(Number);

  const [ordershipment, setOrdershipment] = useState(String);
  const [orderdetail, setOrderdetail] = useState(String);
  //const [datetime, setDatetime] = useState(String);

  let stock = Number(orderstockid)
  let managerID = Number(cookieID)
  let typeproductID = Number(typeproductid)
  let productID = Number(productid)
  let companyID = Number(companyid)

  let detail = String(orderdetail)
  let shipment = String(ordershipment)
  console.log(managerID)
  useEffect(() => {

    /*const getmanagers = async () => {

      const mn = await api.listManager({ limit: 10, offset: 0 });
      setLoading(false);
      setManagers(mn);
    };
    getmanagers();*/

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
    shipment,
    detail,
  }

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'stock':
        alertMessage("error","กรุณาใส่จำนวนให้ถูกต้อง");
        return;
      case 'shipment':
        alertMessage("error","กรุณาระบุการจัดส่งให้ถูกต้อง");
        return;
      case 'detail':
        alertMessage("error","กรุณาหมายเหตุให้ครบถ้วน");
        return;
      default:
        alertMessage("error","<h2>บันทึกข้อมูลไม่สำเร็จ</h2>");
        return;
    }
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
          if (data.status == true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',

          }); window.setTimeout(function () { location.reload() }, 3000);
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
  }


  const Product_id_handleChange = (event: any) => {
    setProductid(event.target.value);
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
                to="/orderproduct">
                Back
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
                ระบบค้นหารายการสั่งซื้อสินค้าเข้ามาในคลัง
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
          <Tab textColor="inherit" label="Search Data" />
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
                ชื่อผู้ใช้ระบบค้นหา
              </Typography>
            </Grid>
            <Grid item xs={2}>

              <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ชื่อสินค้าที่ต้องการค้นหา
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
      <ComponanceTable></ComponanceTable>

    </div>
  );
}