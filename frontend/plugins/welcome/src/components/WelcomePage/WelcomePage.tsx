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
            image="../../image/account.jpg"
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
      <Header style={HeaderCustom} title={`ระบบ...`}></Header>
      <Content>
        <ContentHeader title="สมาชิกในกลุ่ม"></ContentHeader>
        <Grid container>
          <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
          <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
          <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
          <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
          <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
          <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
          
          
          <Button
                  style={{ marginLeft: 1 }}
                  component={RouterLink}
                  to="/SignInOrderonline"
                  variant="contained"
                >
                  Back
             </Button>

             <Button
                  style={{ marginLeft: 1 }}
                  component={RouterLink}
                  to="/Stock"
                  variant="contained"
                >
                  Back2
             </Button>
             <Button
                  style={{ marginLeft: 1 }}
                  component={RouterLink}
                  to="/Salary"
                  variant="contained"
                >
                  Back to Salary
             </Button>
             <Button
                  style={{ marginLeft: 1 }}
                  component={RouterLink}
                  to="/Promotion"
                  variant="contained"
                >
                  Back to Promotion
             </Button>
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
