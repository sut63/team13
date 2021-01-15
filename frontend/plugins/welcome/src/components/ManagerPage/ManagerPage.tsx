import React, { FC } from 'react';
import { Typography, Grid, Button } from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';

const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles({
  root: {
    maxWidth: 345,
  },
});

export type ProfileProps = {
  name: string; 
  id: string;
  system: string;
};

export function CardTeam({ name, id, system }: ProfileProps) {
  const classes = useStyles();
  return (
    <Grid item xs={12} md={3}>
      <Card className={classes.root}>
        <CardActionArea>
          <CardMedia
            component="img"
            alt="นาย สมชาย ใจดี"
            height="140"
            image="E:\SE\Project\Picture\1.png"
            title="นาย สมชาย ใจดี"
          />
          <CardContent>
            <Typography gutterBottom variant="h5" component="h2">
              {system}
            </Typography>
            <Typography gutterBottom variant="h5" component="h2">
              {id} {name}
            </Typography>
          </CardContent>
        </CardActionArea>
      </Card>
    </Grid>
  );
}

const WelcomePage: FC<{}> = () => {
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบ Farm Mart`}></Header>
      <Content>
        <ContentHeader title="สมาชิกในกลุ่ม"></ContentHeader>
        <Grid container>
          <CardTeam name={"นาย นนทกร พาอยู่สุข"} id={"B6111052"} system={"ระบบย่อย: ระบบบันทึกเงินเดือนพนักงาน"}></CardTeam>
          <CardTeam name={"นาย ธนบดี เพชรรี่"} id={"B6118631"} system={"ระบบย่อย: ระบบตารางเวลาทำงานพนักงาน"}></CardTeam>
          <CardTeam name={"นาย พิชัย โสมาสา "} id={"B6116637"} system={"ระบบย่อย: ระบบPromotion "}></CardTeam>
          <CardTeam name={"นาย ภูมิมินทร์ พินพิมาย"} id={"B6111090"} system={"ระบบย่อย: ระบบสั่งซื้อสินค้าเข้ามาในคลัง"}></CardTeam>
        
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
