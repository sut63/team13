
import {
  Grid,
  MenuItem,
  
  TextField,
  Button,
  FormControl,
  InputLabel,
  Select,
  Box,
  Avatar,
} from '@material-ui/core';
import React, { FC,useEffect, useState } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Menu from '@material-ui/core/Menu';
import { DefaultApi } from '../../api/apis';


import { Link as RouterLink } from 'react-router-dom';


import { Alert } from '@material-ui/lab';


import { EntProduct } from '../../api/models/EntProduct';
import { EntTypeproduct } from '../../api/models/EntTypeproduct';
import { EntEmployee} from '../../api/models/EntEmployee';
import { EntZoneproduct} from '../../api/models/EntZoneproduct';
import { Content, ContentHeader, Header, Page, pageTheme } from '@backstage/core';


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

const Stock: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'Your Stock' };






  const api = new DefaultApi();
  const [products, setProducts] = useState<EntProduct[]>([]);
  const [typeproducts, setTypeproducts] = useState<EntTypeproduct[]>([]);
  const [employees, setEmployees] = useState<EntEmployee[]>([]);
  const [zoneproducts, setZoneproducts] = useState<EntZoneproduct[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [productid, setProductid] = useState(Number);
  const [priceproduct, setPriceproduct] = useState(Number);
  const [amount, setAmount] = useState(Number);
  const [time, setTime] = useState(String);
  const [typeproductid, setTypeproductid] = useState(Number);
  const [employeeid, setEmployeeid] = useState(Number);
  const [zoneproductid, setZoneproductid] = useState(Number);

  useEffect(() => {
    const getProducts = async () => {
      const pr = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(pr);
    };
    getProducts();


    const getTypeproducts = async () => {
      const ty = await api.listTypeproduct({ limit: 10, offset: 0 });
      setLoading(false);
      setTypeproducts(ty);
    };
    getTypeproducts();

    const getEmployees = async () => {
      const em = await api.listEmployee({ limit: 10, offset: 0 });
      setLoading(false);
      setEmployees(em);
    };
    getEmployees();


    const getZoneproducts = async () => {
      const zo = await api.listZoneproduct({ limit: 10, offset: 0 });
      setLoading(false);
      setZoneproducts(zo);
    };
    getZoneproducts();

  }, [loading]);

  const handletimeChange = (event: any) => {
      setTime(event.target.value as string);
    };

  const CreateStock = async () => {
    const stock = {
      product: productid,
      typeproduct: typeproductid,
      employee: employeeid,
      zoneproduct: zoneproductid,
      priceproduct : priceproduct,
      amount : amount,
      time: time + "00:00+07:00"



    }
    console.log(stock);

    const res: any = await api.createStock({ stock: stock });
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
    setProductid(event.target.value as number);
  };
  const employee_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmployeeid(event.target.value as number);
  };

  const typeproduct_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setTypeproductid(event.target.value as number);
  };

  const zoneproduct_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setZoneproductid(event.target.value as number);
    };

  const amount_id_handleChange = (event: any) => {
      setAmount(event.target.value);
     };

  const priceproduct_id_handleChange = (event: any) => {
      setPriceproduct(event.target.value);
     };



  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to Stock Product'}`}
        subtitle="Product in your stock."
      >

        <Avatar>D</Avatar>
        <Typography component="div" variant="body1">
          <Box color="Pang@gmail.com">Pang@gmail.com</Box>
          <Box color="secondary.main"></Box>
        </Typography>

      </Header>
      <Content>

        <ContentHeader title="Your Stock">          
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
          <form noValidate autoComplete="off"></form>



    
      

      
        <Grid container alignItems="center" spacing={4}>
          <Grid item xs={12}></Grid>
          <Grid item xs={12}></Grid>




          <Grid item xs={4}>   <center>
            <h2 align='right'>Employee</h2></center>
          </Grid>
          <Grid item xs={4}>
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            ><InputLabel>Employee</InputLabel>
              <Select
                labelId="employee_id-label"
                label="employee"
                id="eployee_id"
                value={employeeid}
                onChange={employee_id_handleChange}
                style={{ marginRight : 300 ,width: 300 }}
              >
                {employees.map((item: EntEmployee) =>
                  <MenuItem value={item.id}>{item.name}</MenuItem>)}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={4}>   
          </Grid>





          <Grid item xs={4}> <center>
            <h2 align='right'>Typeproduct</h2></center>

          </Grid>
          <Grid item xs={4}>
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            ><InputLabel>Typeproduct</InputLabel>
              <Select
                labelId="typeproduct_id-label"
                label="typeproduct"
                id="typeproduct_id"
                value={typeproductid}
                onChange={typeproduct_id_handleChange}
                style={{marginRight:300 , width: 300 }}
              >
                {typeproducts.map((item: EntTypeproduct) =>
                  <MenuItem value={item.id}>{item.typeproduct}</MenuItem>)}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={4}>   
          </Grid>






          <Grid item xs={4}><center>
            <h2 align='right'>Product</h2></center>
          </Grid>
          <Grid item xs={4}>
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            ><InputLabel>Product</InputLabel>
              <Select
                labelId="product_id-label"
                label="product"
                id="product_id"
                value={productid}
                onChange={product_id_handleChange}
                style={{ marginRight:300, width: 300 }}
              >
                {products.map((item: EntProduct) =>
                  <MenuItem value={item.id}>{item.nameProduct}</MenuItem>)}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={4}>   
          </Grid>




          <Grid item xs={4}><center>
            <h2 align='right'>Zoneroduct</h2></center>
          </Grid>
          <Grid item xs={4}>
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            ><InputLabel>Zoneproduct</InputLabel>
              <Select
                labelId="zoneproduct_id-label"
                label="zoneproduct"
                id="zoneproduct_id"
                value={zoneproductid}
                onChange={zoneproduct_id_handleChange}
                style={{marginRight:300, width: 300 }}
              >
                {zoneproducts.map((item: EntZoneproduct) =>
                  <MenuItem value={item.id}>{item.zone}</MenuItem>)}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={4}>   
          </Grid>





          <Grid item xs={4}><center>
            <h2 align='right'>Price</h2></center>
          </Grid>
          <Grid item xs={4}>
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"

            ><TextField
                id="price"
                label="Price"
                variant="outlined"
                type="string"
                size="medium"
                value={priceproduct}
                onChange={priceproduct_id_handleChange }
                style={{  marginRight :300,width: 300 }}
              />
            </FormControl>
          </Grid>
          <Grid item xs={4}>   
          </Grid>


          
          <Grid item xs={4}><center>
            <h2 align='right'>Amount</h2></center>
          </Grid>
          <Grid item xs={4}>
          <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
              style={{ marginRight :300 ,width: 300 }}
            >
              <TextField id="outlined-number" type='number'  InputLabelProps={{
                shrink: true,}}label="Amount" variant="outlined"
                onChange = {amount_id_handleChange}
                />
            </FormControl>
          </Grid>
          <Grid item xs={4}>   
          </Grid>




          

          <Grid item xs={4}><center>
            <h2 align='right'>Date</h2></center>
          </Grid>
          <Grid item xs={4}>
          
        <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
              style={{ marginRight: 500, width: 2000 }}
            >
              <TextField
                id="Date"
                label="Date"
                type="date"
                value={time}
                onChange={handletimeChange}
                //defaultValue="2020-05-24"
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
              />
            </FormControl>
            
            </Grid>
          <Grid item xs={4}>   
          </Grid>







          <Grid item xs={4}>
          </Grid>
          <Grid item xs={4}>   

            <Button
              onClick={() => {
                CreateStock();
              }}
              variant="contained"
              color="primary"
              style={{ marginLeft: 20 ,width : 100 }}
            >
              Save
             </Button>
            <Button
              style={{ marginLeft: 20 ,width : 100 }}
              component={RouterLink}
              to="/login"
              variant="contained"
            >
              Show
             </Button>
               
        

          </Grid>
          
          <Grid item xs={4}></Grid>
          <Grid item xs={12}></Grid>
          <Grid item xs={12}></Grid>
        </Grid>
     
        </div>
        
        
      </Content>
    </Page>
  );
}
export default Stock;



