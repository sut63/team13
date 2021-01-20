
import {
  Grid,
  MenuItem,

  TextField,
  Button,
  FormControl,
  InputLabel,
  Select,
 
} from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import IconButton from '@material-ui/core/IconButton';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Menu from '@material-ui/core/Menu';
import { DefaultApi } from '../../api/apis';
import SaveIcon from '@material-ui/icons/Save';
import { Link as RouterLink } from 'react-router-dom';
import { Alert } from '@material-ui/lab';
import { EntProduct } from '../../api/models/EntProduct';
import { EntTypeproduct } from '../../api/models/EntTypeproduct';
import { EntEmployee } from '../../api/models/EntEmployee';
import { EntZoneproduct } from '../../api/models/EntZoneproduct';
import { Content, ContentHeader, Header, Page, pageTheme } from '@backstage/core';
import Swal from 'sweetalert2';
import { Cookies } from '../Stock/LoginEmployee/Cookie';


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

var ck = new Cookies()
var cookieEmail = ck.GetCookie()
var cookieID = ck.GetID()
var cookieName = ck.GetName()

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
  employeeID: number;
  typeproductID: number;
  productID: number;
  zoneproductID: number;
  amount: number;
  priceproduct: number;
  time: Date;
}


export default function Stock() {
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
  const [order, setOreder] = React.useState<Partial<order>>({});
  const [productid, setProductid] = useState(Number);
  const [priceproducts, setPriceproduct] = useState(Number);
  const [amounts, setAmount] = useState(Number);
  const [time, setTime] = useState(String);
  const [typeproductid, setTypeproductid] = useState(Number);
  const [employeeid, setEmployeeid] = useState(Number);
  const [zoneproductid, setZoneproductid] = useState(Number);
  const [idcardemployee, setidcardemployee] = useState(String);

  const [IdcardemployeeError, setIdcardemployeeError] = React.useState('');
  const [PriceError, setPriceError] = React.useState('');
  const [AmountError, setAmountError] = React.useState('');
  const [errors, setError] = React.useState(String);

  let productID = Number(productid)
  let employeeID = Number(cookieID)
  let zoneID = Number(zoneproductid)
  let typeproductID = Number(typeproductid)
  let amount = Number(amounts)
  let priceproduct = Number(priceproducts)
  let iDcardemployee = String(idcardemployee)





 // console.log(employeeID)
  useEffect(() => {

    const getEmployees = async () => {
      const em = await api.listEmployee({ limit: 10, offset: 0 });
      setLoading(false);
      setEmployees(em);
    };
    getEmployees();


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




    const getZoneproducts = async () => {
      const zo = await api.listZoneproduct({ limit: 10, offset: 0 });
      setLoading(false);
      setZoneproducts(zo);
    };
    getZoneproducts();

  }, [loading]);



  const stock = {

    employeeID,
    typeproductID,
    productID,
    zoneID,
    priceproduct,
    iDcardemployee,
    amount,
    time: time + ":00+07:00"
  }
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }


 /* //validate
  const validateIdcardemployee = (val: string) => {
    return val.match("[E]\\d{4}");
  }

  const validateAmount = (val: Number) => {
    return val > 0  ? true : false;
  }

  // ฟังก์ชั่นสำหรับ validate 
  const validatePrice = (val: Number) => {
    return val  > 0 ? false : true;
  }
*/

  /* checkPattern
  const checkPattern  = (id: string, value: string) => {
    switch(id) { 
      case 'IDcardemployee':
        validateIdcardemployee(value) ? setIdcardemployeeError('') : setIdcardemployeeError('ห้ามเกิน 50 ตัวอักษร');
        return;
      case 'Priceproduct': 
        validatePrice(Number(value)) ? setPriceError('') : setPriceError('Ex 0850583300');
        return;
      case 'Amount':
        validateAmount(Number(value)) ? setAmountError('') : setAmountError('เข้าพักได้ไม่เกินห้องละ 5 คน')
        return;
      default:
        return;
    }
  }*/

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'IDcardemployee':
        alertMessage("error","รูปแบบรหัสพนักงานไม่ถูกต้อง กรุณากรอกข้อมูลให้ถูกต้อง");
        return;
      case 'Priceproduct':
        alertMessage("error","ราคาต้องเป็นตัวเลขและห้ามติดลบ กรุณากรอกราคาให้ถูกต้อง");
        return;
      case 'Amount':
        alertMessage("error","จำนวนต้องเป็นตัวเลขห้ามติดลบ กรุณากรอกจำนวนให้ถูกต้อง");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }




  console.log(stock)

  function CreateStock() {
    const apiUrl = 'http://localhost:8080/api/v1/stocks';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(stock),
    };

    console.log(stock);

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
  };

  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);


  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
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
    //checkPattern(data.error.Name);
  };

  const idcardemployee_id_handleChange = (event: any) => {
    setidcardemployee(event.target.value);
    //checkPattern(data.error.Name);

  };

  const handletimeChange = (event: any) => {
    setTime(event.target.value as string);
  };


  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to Stock Product'}`}
        subtitle="Add Product in your stock."

        
      >
        <div style={{ marginLeft: 300 }}>{cookieName}</div>
   <IconButton
                aria-label="account of current patientofphysician"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={handleMenu}
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
                <Button className={classes.button} variant="outlined" color="inherit"
                  size="small" component={RouterLink}
                  to="/LoginEmployee">
                  logout
              </Button>
              </Menu>


      </Header>
      <Content>

        <ContentHeader title="Add Your Product here">
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
              <h3 align='right'>ชื่อพนักงาน</h3></center>
            </Grid>
            <Grid item xs={4}>
              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
              ><InputLabel></InputLabel>

                {/*} <Select
                labelId="employee_id-label"
                label="employee"
                id="eployee_id"
                value={employeeid}
                onChange={employee_id_handleChange}
                style={{ marginRight : 300 ,width: 300 }}
              >
                {employees.map((item: EntEmployee) =>
                  <MenuItem value={item.id}>{item.name}</MenuItem>)}
                </Select>*/}
                <div style={{ marginRight: 300 }}>{cookieName}</div>

              </FormControl>
            </Grid>
            <Grid item xs={4}>
            </Grid>




            <Grid item xs={4}><center>
              <h3 align='right'>รหัสพนักงาน</h3></center>
            </Grid>
            <Grid item xs={4}>
              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"

              ><TextField
                  id="Idcardemployee"
                  label="Idcardemployee"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={idcardemployee}
                  onChange={idcardemployee_id_handleChange}
                  style={{ marginRight: 300, width: 300 }}
                />
              </FormControl>
            </Grid>
            <Grid item xs={4}>
            </Grid>



            <Grid item xs={4}> <center>
              <h3 align='right'>ประเภทสินค้า</h3></center>

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
                  style={{ marginRight: 300, width: 300 }}
                >
                  {typeproducts.map((item: EntTypeproduct) =>
                    <MenuItem value={item.id}>{item.typeproduct}</MenuItem>)}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={4}>
            </Grid>






            <Grid item xs={4}><center>
              <h3 align='right'>สินค้า</h3></center>
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
                  style={{ marginRight: 300, width: 300 }}
                >
                  {products.map((item: EntProduct) =>
                    <MenuItem value={item.id}>{item.nameProduct}</MenuItem>)}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={4}>
            </Grid>




            <Grid item xs={4}><center>
              <h3 align='right'>ตำแหน่งที่วางสินค้า</h3></center>
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
                  style={{ marginRight: 300, width: 300 }}
                >
                  {zoneproducts.map((item: EntZoneproduct) =>
                    <MenuItem value={item.id}>{item.zone}</MenuItem>)}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={4}>
            </Grid>





            <Grid item xs={4}><center>
              <h3 align='right'>ราคาสินค้า</h3></center>
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
                  value={priceproducts}
                  onChange={priceproduct_id_handleChange}
                  style={{ marginRight: 300, width: 300 }}
                />
              </FormControl>
            </Grid>
            <Grid item xs={4}>
            </Grid>



            <Grid item xs={4}><center>
              <h3 align='right'>จำนวนสินค้า</h3></center>
            </Grid>
            <Grid item xs={4}>
              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginRight: 300, width: 300 }}
              >
                <TextField id="outlined-number" type='number' InputLabelProps={{
                  shrink: true,
                }} label="Amount" variant="outlined"
                  onChange={amount_id_handleChange}
                />
              </FormControl>
            </Grid>
            <Grid item xs={4}>
            </Grid>






            <Grid item xs={4}><center>
              <h3 align='right'>เวลาที่บันทึก</h3></center>
            </Grid>
            <Grid item xs={4}>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginRight: 500, width: 2000 }}
              >
                <TextField
                  id="date"
                  label="Datetime"
                  type="datetime-local"
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
        size="large"
        className={classes.button}
        startIcon={<SaveIcon />}
      >
        Save
      </Button>
              
             <Button variant="contained" size="large" color="primary" className={classes.margin}
             component={RouterLink}
             to="/Tablestock">
          Showstock
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
      




