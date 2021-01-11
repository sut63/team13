import React, { FC, useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import { InputLabel, MenuItem, Select } from '@material-ui/core';
import Typography from '@material-ui/core/Typography';
import TableCell from '@material-ui/core/TableCell';
import Avatar from '@material-ui/core/Avatar';
import Box from '@material-ui/core/Box';

import { EntProduct } from '../../api/models/EntProduct';
import { EntTypeproduct } from '../../api/models/EntTypeproduct';
import { EntPaymentchannel } from '../../api/models/EntPaymentchannel';
import { EntCustomer } from '../../api/models/EntCustomer';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(2),
    },
    withoutLabel: {
      marginTop: theme.spacing(2),
    },
    textField: {
      width: '25ch',
    },
  }),
);

const OrderOnline: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'to Order Online' };
  const api = new DefaultApi();

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [products, setProducts] = React.useState<EntProduct[]>([]);
  const [typeproducts, setTypeproducts] = React.useState<EntTypeproduct[]>([]);
  const [paymentchannels, setPaymentchannels] = React.useState<EntPaymentchannel[]>([]);
  const [customers, setCustomers] = React.useState<EntCustomer[]>([]);

  const [product, setProduct] = useState(Number);
  const [typeduct, setTypeproduct] = useState(Number);
  const [paymentchannel, setPaymentchannel] = useState(Number);
  const [customer, setCustomer] = useState(Number);
  const [stock, setScotk] = useState(Number);
  const [addedtime, Setaddedtime] = useState(String);


  useEffect(() => {
    const getProducts = async () => {
      const p = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(p);
    };
    getProducts();

    const getTypeproducts = async () => {
      const d = await api.listTypeproduct({ limit: 10, offset: 0 });
      setLoading(false);
      setTypeproducts(d);
    };
    getTypeproducts();

    const getPaymentchannels = async () => {
      const pay = await api.listPaymentchannel({ limit: 10, offset: 0 });
      setLoading(false);
      setPaymentchannels(pay);
    };
    getPaymentchannels();

    const getCustomers = async () => {
        const c = await api.listCustomer({ limit: 10, offset: 0 });
        setLoading(false);
        setCustomers(c);
      };
      getCustomers();

  }, [loading]);


  const handletimeChange = (event: any) => {
    Setaddedtime(event.target.value as string);
  };

  const CreateOrderonline = async () => {
    const orderonline = {
      productid: product,
      typeductid: typeduct,
      paymentchannelid: paymentchannel,
      customerid: customer,
      stock: stock,
      addedtime: addedtime + "00:00+07:00"
    }

    const res: any = await api.createOrderonline({ orderonline: orderonline });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };


  const product_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProduct(event.target.value as number);
  };

  const typeproduct_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setTypeproduct(event.target.value as number);
  };

  const paymentchannel_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPaymentchannel(event.target.value as number);
  };

  const customer_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setCustomer(event.target.value as number);
  };

  const Stock_id_handleChange = (event: any) => {
    setScotk(event.target.value);
   };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to Orderonline'}`}
        subtitle="Select Product you want to be in."
      >

        <Avatar>D</Avatar>
        <Typography component="div" variant="body1">
          <Box color="Dang@gmail.com">Dang@gmail.com</Box>
          <Box color="secondary.main"></Box>
        </Typography>

      </Header>
      <Content>

        <ContentHeader title="Orderonline">          
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  This is a success alert — check it out!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>

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
                  id="DayStart"
                  label="DayStart"
                  type="date"
                  value={addedtime}
                  onChange={handletimeChange}
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
                  value={product}
                  onChange={product_id_handleChange}
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
                  value={typeduct}
                  onChange={typeproduct_id_handleChange}
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
                  value={paymentchannel}
                  onChange={paymentchannel_id_handleChange}
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
                  value={customer}
                  onChange={customer_id_handleChange}
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
                <TextField id="outlined-number" type='number'  InputLabelProps={{
                  shrink: true,}}label="กรุณาใส่จำนวน" variant="outlined"
                  onChange = {Stock_id_handleChange}
                  />
              </FormControl>


            </TableCell>

            <div className={classes.margin}>
              <TableCell align="right">
                <Button
                  onClick={() => {
                    CreateOrderonline();
                  }}
                  variant="contained"
                  color="primary"
                  style={{ marginLeft: 545, width: 200 }}
                >
                  SAVE DATA
             </Button>
              </TableCell>

              <TableCell align="right">
                <Button
                  style={{ marginLeft: 1 }}
                  component={RouterLink}
                  to="/WelcomePage"
                  variant="contained"
                >
                  Back
             </Button>
              </TableCell>
              
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
};

export default OrderOnline;
