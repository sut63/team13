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
import Swal from 'sweetalert2';
import SearchIcon from '@material-ui/icons/Search';
import { EntPromotion } from '../../../api';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
import { Cookies } from '../../orderproduct/SignInOrderproduct/Cookie';

const lightColor = 'rgba(255, 255, 255, 0.7)';
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
  const api = new DefaultApi();

  const [products, setProducts] = useState<EntProduct[]>([]);
  const [loading, setLoading] = useState(true);
  const [productid, setProductid] = useState(Number);

  let managerID = Number(cookieID)
  let productID = Number(productid)
  console.log(managerID)

  const [promotion, setPromotion] = useState<EntPromotion[]>();


  useEffect(() => {
    const getproducts = async () => {
      const pr = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(pr);
    };
    getproducts();
  }, [loading]);

  const Promotion = {
    managerID,
    productID,
  }
  console.log(Promotion)

  const Product_id_handleChange = (event: any) => {
    setProductid(event.target.value);
  }
  var lenPromotion: number
  
  const getCheckinsorder = async () => {
    const res = await api.getPromotion({ id: productid })
    setPromotion(res)
    lenPromotion = res.length
    if (lenPromotion > 0) {
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
                to="/Promotiontable">
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
                ระบบค้นหาโปรโมชั่นของสินค้า
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
      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
            <TableCell align="center">No.</TableCell>
              <TableCell align="center">ชื่อโปรโมชั่น</TableCell>
              <TableCell align="center">สินค้า</TableCell>
              <TableCell align="center">Barcode</TableCell>
              <TableCell align="center">วันหมดอายุ</TableCell>
              <TableCell align="center">ของแถม</TableCell>
              <TableCell align="center">ส่วนลด</TableCell>
              <TableCell align="center">ราคา</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {promotion === undefined
              ? null
              : promotion.map((item: any) => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.promotionName}</TableCell>
                  <TableCell align="center">{item.edges?.product?.nameProduct}</TableCell>
                  <TableCell align="center">{item.edges?.product?.barcodeProduct}</TableCell>
                  <TableCell align="center">{item.edges?.product?.eXP}</TableCell>
                  <TableCell align="center">{item.edges?.give?.giveawayName}</TableCell>
                  <TableCell align="center">{item.edges?.sale?.sale}</TableCell>
                  <TableCell align="center">{item.price}</TableCell>
                  <TableCell align="center">
                  </TableCell>
                </TableRow>
              ))}
          </TableBody>
        </Table>
      </TableContainer>

    </div>
  );
}
