import React, { useEffect, useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Grid from '@material-ui/core/Grid';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import Button from '@material-ui/core/Button';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import { Paper, Select, TextField } from '@material-ui/core';
import { ContentHeader } from '@backstage/core';
import Swal from 'sweetalert2';

import { EntProduct } from '../../api/models/EntProduct';
import { EntDiscount } from '../../api/models/EntDiscount';
import { EntGiveaway } from '../../api/models/EntGiveaway';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
  }),
);

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

export default function MenuAppBar() {
  const classes = useStyles();
  const [auth, setAuth] = React.useState(true);
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);
  const api = new DefaultApi();

  const [status, setStatus] = useState(false);
  const [loading, setLoading] = useState(true);

  const [products, setProducts] = React.useState<EntProduct[]>([]);
  const [discount, setDiscount] = React.useState<EntDiscount[]>([]);
  const [giveaway, setGiveaway] = React.useState<EntGiveaway[]>([]);


  const [productID, setProductid] = useState(Number);
  const [discountID, setDiscountid] = useState(Number);
  const [giveawayID, setGiveawayid] = useState(Number);
  const [promotionname, setPromotionname] = useState(String);
  const [durationpromotion, setDurationpromotion ] = useState(String);
  const [price, setPrice] = useState(Number);



  // Lifecycle Hooks
  useEffect(() => {

    const getDiscount = async () => {
      const dc = await api.listDiscount({ limit: 10, offset: 0 });
      setDiscount(dc);
    };
    getDiscount();

    const getGiveaway = async () => {
      const gw = await api.listGiveaway({ limit: 10, offset: 0 });
      setGiveaway(gw);
    };
    getGiveaway();

    const getproducts = async () => {
      const pr = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(pr);
    };
    getproducts();

  }, [loading]);

  let pc = Number(price)
  
    const Promotions = {
      discount: discountID,
      product: productID,
      giveaway: giveawayID,
      DurationPromotion: durationpromotion,
      promotionName: promotionname,
      price: pc,
    }


    const alertMessage = (icon: any, title: any) => {
      Toast.fire({
        icon: icon,
        title: title,
      });
    }

    const checkCaseSaveError = (field: string) => {
      switch(field) {
        case 'PromotionName':
          alertMessage("error","กรุณาเขียนชื่อ promotion เพิ่มเติม");
          return;
        case 'Price':
          alertMessage("error","กรุณากรอกราคาให้ถูกต้อง");
          return;
        case 'DurationPromotion':
          alertMessage("error","กรุณาระบุระยะเวลาโปรโมชั่นให้ถูกต้อง");
          return;
        default:
          alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
          return;
      }
    }
    
console.log(Promotions)
function save() {
  const apiUrl = 'http://localhost:8080/api/v1/promotions';
  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(Promotions),
  };

  console.log(Promotions); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

  fetch(apiUrl, requestOptions)
    .then(response => response.json())
    .then(data => {
      console.log(data);
      if (data.status == true) {
        //clear();
        Toast.fire({
          icon: 'success',
          title: 'บันทึกข้อมูลสำเร็จ',

        });//window.setTimeout(function(){location.reload()},8000);
      } else {
        checkCaseSaveError(data.error.Name)
      }
    });

    
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  const Product_id_handleChange = (event: any) => {
    setProductid(event.target.value);
  };

  const discount_id_handleChange = (event: any) => {
    setDiscountid(event.target.value);
  };

  const giveaway_id_handleChange = (event: any) => {
    setGiveawayid(event.target.value);
  };

  const price_id_handleChange = (event: any) => {
    setPrice(event.target.value as number);
  };

  const Promotionname_id_handleChange = (event: any) => {
    setPromotionname(event.target.value as string);
  };

  const Duration_id_handleChange = (event: any) => {
    setDurationpromotion(event.target.value as string);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h5" className={classes.title}>
            Create PROMOTION
          </Typography>
          {auth && (
            <div>
              <IconButton
                aria-label="account of current user"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={handleMenu}
                color="inherit"
              >
                <AccountCircle />
              </IconButton>
              <Menu
                id="menu-appbar"
                anchorEl={anchorEl}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                open={open}
                onClose={handleClose}
              >
                <MenuItem onClick={handleClose}>Logout</MenuItem>
              </Menu>
            </div>
          )}
        </Toolbar>
      </AppBar>

      <AppBar position="static" color='inherit' background-color="inherit">
        <Grid container alignItems="center" spacing={2} >
          <Grid item xs={12}></Grid>
          <Grid item xs={12}></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><h3>กรุณากรอกชื่อโปรโมชั่น</h3></Grid>
          <Grid item xs={2}>
          <TextField id="outlined-number" type='string' InputLabelProps={{
                  shrink: true,
                }} label="" variant="outlined"
                  onChange={Promotionname_id_handleChange}
                  value={promotionname}
                  
          />
          </Grid>
          <Grid item xs={2}><p>(ชื่อ Promotion มากกว่า 10 ตัวอักษร)</p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><h3>เลือกสินค้า</h3></Grid>
          <Grid item xs={2}>
          <Select
              labelId="Product"
              label="Product"
              id="product_id"
              value={productID}
              onChange={Product_id_handleChange}
              style={{ width: 200 }}
            >
              {products.map((item: EntProduct) =>
                <MenuItem key={item.id} value={item.id}>{item.nameProduct}</MenuItem>)}
            </Select>
          </Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><h3>โปรดเลือกของแถม</h3></Grid>
          <Grid item xs={2}>
          <Select
              labelId="Giveaway"
              label="Giveaway"
              id="Giveaway_id"
              value={giveawayID}
              onChange={giveaway_id_handleChange}
              style={{ width: 200 }}
            >
              {giveaway.map((item: EntGiveaway) =>
                <MenuItem key={item.id} value={item.id}>{item.giveawayName}</MenuItem>)}
            </Select>

          </Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><h3>โปรดเลือกส่วนลด</h3></Grid>
          <Grid item xs={2}>
          <Select
              labelId="Discount"
              label="Discount"
              id="Discount_id"
              value={discountID}
              onChange={discount_id_handleChange}
              style={{ width: 200 }}
            >
              {discount.map((item: EntDiscount) =>
                <MenuItem key={item.id} value={item.id}>{item.sale}</MenuItem>)}
            </Select>
          </Grid>
          <Grid item xs={2}><p>%</p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><h3>ระยะเวลาโปรโมชั่น</h3></Grid>
          <Grid item xs={2}>
          <Paper >
                <TextField id="ระยะ Promotion" type='string'  InputLabelProps={{
                  shrink: true,}}label="" variant="outlined"
                  onChange = {Duration_id_handleChange}
                  value={durationpromotion}
                  />
          </Paper>
          </Grid>
          <Grid item xs={2}><p>(วันที่เริ่มโปรโมชั่น - วันที่หมดโปรโมชั่น)(ตัวอย่าง: 3 มค - 10 มค 2564)</p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><h3>กรุณาระบุราคา</h3></Grid>
          <Grid item xs={2}>          
          <Paper >
                <TextField id="outlined-number" type='number'  InputLabelProps={{
                  shrink: true,}}label="" variant="outlined"
                  onChange = {price_id_handleChange}
                  value={price}
                  />
          </Paper></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={3}><p></p></Grid>
          <Grid item xs={3}></Grid>
          <Grid item xs={3}></Grid>
          <Grid item xs={3}><p></p></Grid>

          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}>
          <Button
              style={{ marginLeft: 30 }}
              component={RouterLink}
              to="/Promotiontable"
              variant="contained"
            >
              show Promotion
            </Button>
            
          </Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}>
            <Button
              onClick={() => {
                save();
              }}

              color="inherit"
              style={{ width: 200 }}
              variant="contained">

              SAVE
            </Button>
          </Grid>
          <Grid item xs={1}></Grid>
          <Grid item xs={1}>
            <Button
              style={{ marginLeft: 20 }}
              component={RouterLink}
              to="/SplitsystemManager"
              variant="contained"
            >
              Back
            </Button>
          </Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>



          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

        </Grid>
      </AppBar>


      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
          </Typography>
        </Toolbar>
      </AppBar>
    </div>



  );
}