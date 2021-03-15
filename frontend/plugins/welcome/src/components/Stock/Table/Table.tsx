import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import DeleteIcon from '@material-ui/icons/Delete';
import { Link as RouterLink } from 'react-router-dom';
import { AppBar, IconButton, Menu, Toolbar, Typography } from '@material-ui/core';
import {
  MenuItem, 
  Button, 
} from '@material-ui/core';
import AccountCircle from '@material-ui/icons/AccountCircle';
//import { DefaultApi } from '../../api/apis';
import {DefaultApi} from '../../../api/apis';
import { EntStock } from '../../../api/models/EntStock';

 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();

 const [stocks, setStocks] = useState<EntStock[]>();
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getStocks = async () => {
     const res = await api.listStock({ limit: 10, offset: 0 });
     setLoading(false);
     setStocks(res);
   };
   getStocks();
 }, [loading]);
 
 const deleteStocks = async (id: number) => {
   const res = await api.deleteStock({ id: id });
   setLoading(true);
 };
 
 return (
 
   
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">Employee</TableCell>
           <TableCell align="center">IDStock</TableCell>
           <TableCell align="center">Typeproduct</TableCell>
           <TableCell align="center">Product</TableCell>
           <TableCell align="center">Zoneproduct</TableCell>
           <TableCell align="center">Priceproduct</TableCell>
           <TableCell align="center">Amount</TableCell>
           <TableCell align="center">Timesave</TableCell>
           <Button
              style={{ marginLeft: 20 ,width : 100 }}
              component={RouterLink}
              to="/Stock"
              variant="outlined"
            >
              Back
             </Button>
         </TableRow>
       </TableHead>
       <TableBody>
       {stocks === undefined 
          ? null
          : stocks.map((item : any)=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.employee?.name}</TableCell>
             <TableCell align="center">{item.iDstock}</TableCell>
             <TableCell align="center">{item.edges?.typeproduct?.typeproduct}</TableCell>
             <TableCell align="center">{item.edges?.product?.nameProduct}</TableCell>
             <TableCell align="center">{item.edges?.zoneproduct?.zone}</TableCell>
             <TableCell align="center">{item.priceproduct}</TableCell>
             <TableCell align="center">{item.amount}</TableCell>
             <TableCell align="center">{item.time}</TableCell>
             
             
             

             <TableCell align="center">

             <Button
              onClick={() => {
                deleteStocks(item.id);
              }}
        variant="outlined"
        color="secondary"
        className={classes.button}
        startIcon={<DeleteIcon />}
      >
        Delete
      </Button>





             </TableCell>
             
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
