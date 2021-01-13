import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  ContentHeader,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';

import {
  makeStyles,
  Theme,
  createStyles,
} from '@material-ui/core/styles';

import {
  TextField,
  Button,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  Typography,
  TableCell,
  Avatar,
  Box,
} from '@material-ui/core';

import { EntProduct } from '../../api/models/EntProduct';
import { EntTypeproduct } from '../../api/models/EntTypeproduct';
import { EntPaymentchannel } from '../../api/models/EntPaymentchannel';
import { EntCustomer } from '../../api/models/EntCustomer';
import Swal from 'sweetalert2';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { DefaultApi } from '../../api/apis';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(1),
    },
    withoutLabel: {
      marginTop: theme.spacing(2),
    },
    textField: {
      width: '25ch',
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

interface order {
  customerid: number;
  typeproductid: number;
  productid: number;
  paymentchannelid: number;
  stock: number;
  addedtime: Date;
}

export default function Orderonline() {
  const classes = useStyles();
  const profile = { givenName: 'to Order Online' };
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);

  const [order, setOreder] = React.useState<Partial<order>>({});

  const [products, setProducts] = React.useState<EntProduct[]>([]);
  const [typeproducts, setTypeproducts] = React.useState<EntTypeproduct[]>([]);
  const [paymentchannels, setPaymentchannels] = React.useState<EntPaymentchannel[]>([]);
  const [customers, setCustomers] = React.useState<EntCustomer[]>([]);

  const [customerID, setCustomerid] = useState(Number);
  const [typeproductID, setTypeproductid] = useState(Number);
  const [productID, setProductid] = useState(Number);
  const [paymentchannelID, setPaymentchannelid] = useState(Number);
  const [orderstockid, setOrderstockid] = useState(Number);
  const [datetime, setDatetime] = useState(String);

  let stock = Number(orderstockid)
  let customerid = Number(customerID)
  let typeproductid = Number(typeproductID)
  let productid = Number(productID)
  let paymentchannelid = Number(paymentchannelID)

  console.log(customerID)
  useEffect(() => {

    const getcustomers = async () => {
      const mn = await api.listCustomer({ limit: 10, offset: 0 });
      setLoading(false);
      setCustomers(mn);
    };
    getcustomers();

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

    const getpaymentchannels = async () => {
      const pay = await api.listPaymentchannel({ limit: 10, offset: 0 });
      setLoading(false);
      setPaymentchannels(pay);
    };
    getpaymentchannels();

  }, [loading]);

  const orderonline = {
    customerid,
    typeproductid,
    productid,
    paymentchannelid,
    stock,
    addedtime: datetime + ":00+07:00"
  }
  console.log(orderonline)

  function clear() {
    Customer_id_handleChange([]);
    Typeproduct_id_handleChange([]);
    Product_id_handleChange([]);
    Paymentchannel_id_handleChange([]);
    Orderstock_id_handleChange([]);
    handleDatetimeChange([]);
  }

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/orderonlines';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(orderonline),    };

    console.log(orderonline); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.id != null) {
          //clear();
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

  const Customer_id_handleChange = (event: any) => {
    setCustomerid(event.target.value);
  };

  const Typeproduct_id_handleChange = (event: any) => {
    setTypeproductid(event.target.value);
  };

  const Product_id_handleChange = (event: any) => {
    setProductid(event.target.value);
  }
  const Paymentchannel_id_handleChange = (event: any) => {
    setPaymentchannelid(event.target.value);
  };
  const Orderstock_id_handleChange = (event: any) => {
    setOrderstockid(event.target.value);
  };
  const handleDatetimeChange = (event: any) => {
    setDatetime(event.target.value as string);
  };


  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to Orderonline'}`}
        subtitle="Select Product you want to be in."
      >

        <Avatar>C</Avatar>
        <Typography component="div" variant="body1">
          <Box color="Dang@gmail.com">Dang@gmail.com</Box>
          <Box color="secondary.main"></Box>
        </Typography>

      </Header>
      <Content>
        <ContentHeader title="PositionAssingment"></ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">


            <TableCell align="left">

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <TextField
                  id="date"
                  label="Date"
                  type="datetime-local"
                  value={datetime || ''}
                  onChange={handleDatetimeChange}
                  //defaultValue="2020-05-24"
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>


              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}

              >
                <InputLabel id="product_id-label">Product</InputLabel>
                <Select
                  labelId="product_id-label"
                  label="Product"
                  id="product_id"
                  value={productID || ''}
                  onChange={Product_id_handleChange}
                  style={{ width: 300 }}
                >
                  {products.map((item: EntProduct) =>
                    <MenuItem value={item.id}>{item.nameProduct}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="typeproduct_id-label">Typeproduct</InputLabel>
                <Select
                  labelId="typeproduct_id-label"
                  label="Typeproduct"
                  id="typeproduct_id"
                  value={typeproductID || ''}
                  onChange={Typeproduct_id_handleChange}
                  style={{ width: 300 }}
                >
                  {typeproducts.map((item: EntTypeproduct) =>
                    <MenuItem value={item.id}>{item.typeproduct}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="paymentchannel_id-label">Paymentchannel</InputLabel>
                <Select
                  labelId="paymentchannel_id-label"
                  label="Paymentchannel"
                  id="paymentchannel_id"
                  value={paymentchannelID || ''}
                  onChange={Paymentchannel_id_handleChange}
                  style={{ width: 300 }}
                >
                  {paymentchannels.map((item: EntPaymentchannel) =>
                    <MenuItem value={item.id}>{item.bank}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="customer_id-label">Customer</InputLabel>
                <Select
                  labelId="customer_id-label"
                  label="Customer"
                  id="customer_id"
                  value={customerID || ''}
                  onChange={Customer_id_handleChange || ''}
                  style={{ width: 300 }}
                >
                  {customers.map((item: EntCustomer) =>
                    <MenuItem value={item.id}>{item.name}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 302 }}
              >
                <TextField id="outlined-number" type='number' InputLabelProps={{
                  shrink: true,
                }} label="กรุณาใส่จำนวน" variant="outlined"
                  onChange={Orderstock_id_handleChange}
                  value={orderstockid || ''}
                />
              </FormControl>


            </TableCell>

            <div className={classes.margin}>
              <TableCell align="right">
                <Button
                  component={RouterLink}
                  to="/Orderonline"
                  variant="contained"
                  color="primary"
                  size="large"
                  style={{ marginLeft: 545, width: 200 }}
                  className={classes.margin}
                  onClick={() => {
                    save();
                  }}
                  startIcon={<SaveIcon
                  />}
                >
                  Save
                </Button>
              </TableCell>

              <TableCell align="right">
                <Button
                  style={{ marginLeft: 1 }}
                  component={RouterLink}
                  to="/"
                  variant="contained"
                >
                  Back
             </Button>
              </TableCell>

              <TableCell align="right">
                <Button
                  style={{ marginLeft: 1 }}
                  component={RouterLink}
                  to="/Orderonlinetable"
                  variant="contained"
                >
                  Show
             </Button>
              </TableCell>

            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}
