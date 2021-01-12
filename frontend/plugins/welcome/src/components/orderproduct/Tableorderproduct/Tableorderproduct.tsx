import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../../api/apis';
import { EntOrderproduct } from '../../../api/models/EntOrderproduct';
 
import moment from 'moment';
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();

 const [orderproducts, setOrderproducts] = useState<EntOrderproduct[]>();
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getOrderproducts = async () => {
     const res = await api.listOrderproduct({ limit: 10, offset: 0 });
     setLoading(false);
     setOrderproducts(res);
   };
   getOrderproducts();
 }, [loading]);
 
 const deleteSystemequipments = async (id: number) => {
   const res = await api.deleteOrderproduct({ id: id });
   setLoading(true);
 };
 console.log(orderproducts)
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
         <TableCell align="center">No.</TableCell>
         <TableCell align="center">Manager</TableCell>
         <TableCell align="center">Product</TableCell>
         <TableCell align="center">Typeproduct</TableCell>
         <TableCell align="center">Company</TableCell>
         <TableCell align="center">Stock</TableCell>
         <TableCell align="center">Date</TableCell>
         <TableCell align="center">Manage</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {orderproducts === undefined 
          ? null
          : orderproducts.map((item :any)=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.manage?.name}</TableCell>
             <TableCell align="center">{item.edges?.product?.NameProduct}</TableCell>
             <TableCell align="center">{item.edges?.typeproduct?.Typeproduct}</TableCell>
             <TableCell align="center">{item.edges?.company?.Name}</TableCell>
             <TableCell align="center">{item.stock}</TableCell>
             <TableCell align="center">{moment(item.addedtime).format('DD/MM/YYYY HH:mm:ss')}</TableCell>
            
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
 );
}
