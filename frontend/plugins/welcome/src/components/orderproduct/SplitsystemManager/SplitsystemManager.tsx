import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import StarIcon from '@material-ui/icons/StarBorder';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';
import { Link as RouterLink } from 'react-router-dom';
import CardMedia from '@material-ui/core/CardMedia';
import CardActionArea from '@material-ui/core/CardActionArea';
import CameraIcon from '@material-ui/icons/PhotoCamera';
import { Cookies } from '../SignInOrderproduct/Cookie'


function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Your Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  '@global': {
    ul: {
      margin: 0,
      padding: 0,
      listStyle: 'none',
    },
  },
  appBar: {
    borderBottom: `1px solid ${theme.palette.divider}`,
  },
  toolbar: {
    flexWrap: 'wrap',
  },
  toolbarTitle: {
    flexGrow: 1,
  },
  link: {
    margin: theme.spacing(1, 1.5),
  },
  heroContent: {
    padding: theme.spacing(8, 0, 6),
  },
  cardHeader: {
    backgroundColor:
      theme.palette.type === 'light' ? theme.palette.grey[200] : theme.palette.grey[700],
  },

  card: {
    height: '200%',
    display: 'flex',
    flexDirection: 'column',
  },
  cardMedia: {
    paddingTop: '100%', // 16:9
  },

  cardPricing: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'baseline',
    marginBottom: theme.spacing(2),
  },
  footer: {
    borderTop: `1px solid ${theme.palette.divider}`,
    marginTop: theme.spacing(8),
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3),
    [theme.breakpoints.up('sm')]: {
      paddingTop: theme.spacing(6),
      paddingBottom: theme.spacing(6),
    },
  },
}));

const tiers = [
  {
    //image: "https://cutt.ly/cjQK0LR",
    image: "https://scontent.fnak3-1.fna.fbcdn.net/v/t1.0-9/110934875_2677631252337858_5542053384630739541_o.jpg?_nc_cat=109&ccb=3&_nc_sid=09cbfe&_nc_eui2=AeG9U6e5y42sY-xZaeT9i1JwOC9uk6eH9fg4L26Tp4f1-ENQz2eKeAhyWiNxNdn09gOMp7UCvoEI7fDFa5r7EA1C&_nc_ohc=QUzRXOPsEL4AX8hNBmL&_nc_ht=scontent.fnak3-1.fna&oh=fbab51a6cde9a51e3d03ea2a9b237b60&oe=604E6433",
    title: 'ระบบบันทึกเงินเดือนพนักงาน',
    description: [/*'ระบบย่อย', 'บันทึกเงินเดือนพนักงาน', 'ตารางเวลาทำงานพนักงาน'*/],
    buttonText: 'Contineus',
    buttinLink: "/Salary",
    buttonVariant: 'contained',
  },
  {
    image: "https://scontent.fnak3-1.fna.fbcdn.net/v/t1.0-9/85148189_1356810611187543_556126289770053632_n.jpg?_nc_cat=107&ccb=3&_nc_sid=174925&_nc_eui2=AeFS7Fb15B3vkYjaR6KPGT9Iz8geMBNEf3XPyB4wE0R_daPxTcDmRu-t9aOUgEPBNQF1fmBivrpsJ0_1LlIZEUIl&_nc_ohc=gmRmUvxh3EoAX_RcRKo&_nc_ht=scontent.fnak3-1.fna&oh=ec152438a620ed4594038cc571d0f89b&oe=604D3DC9",
    title: 'ระบบสั่งซื้อสินค้าเข้ามาในคลัง',
    description: [
     /* '50 users included',
      '30 GB of storage',
      'Help center access',
      'Phone & email support',*/
    ],
    buttonText: 'Contineus',
    buttinLink: "/orderproduct",
    buttonVariant: 'contained',
  },
  {
    image: "https://scontent.fnak3-1.fna.fbcdn.net/v/t1.0-9/61830119_1195335657302819_3015466551059939328_o.jpg?_nc_cat=104&ccb=3&_nc_sid=174925&_nc_eui2=AeFyZveE3nVMZRDpEfNMWEFKMyvJ7xjbOr4zK8nvGNs6vvQohl2u6QNZqtN_2OWGqg03peyBa6tdsetO6lNiUi6s&_nc_ohc=vkA6ZXy5XwEAX8s74fd&_nc_ht=scontent.fnak3-1.fna&oh=62cebe0feb98102f9d63dba0a33ad983&oe=60508EC1",
    title: 'ระบบตารางเวลาทำงานพนักงาน',
    description: [
      /*'20 users included',
      '10 GB of storage',
      'Help center access',
      'Priority email support',*/
    ],
    buttonText: 'Contineus',
    buttinLink: "/EmployeeWorkingHours",
    buttonVariant: 'contained',
  },
  {
    image: "https://scontent.fnak3-1.fna.fbcdn.net/v/t1.0-9/123479591_3473537082733615_3034169409154880874_o.jpg?_nc_cat=109&ccb=3&_nc_sid=09cbfe&_nc_eui2=AeGf4xz8E5Wqmot1Iu_6wc8pWBhQ35dQHyNYGFDfl1AfIyISz0X_lYMq-mS9XiBNYWsmTgq0cH3lPaBx8MGuJSKh&_nc_ohc=lsN8TvtU6TYAX821Tvv&_nc_ht=scontent.fnak3-1.fna&oh=449565778f27f838d51f1fc9b71a6a15&oe=604D0ACF",
    title: 'ระบบPromotion',

    description: [
     /* '50 users included',
      '30 GB of storage',
      'Help center access',
      'Phone & email support',*/
    ],
    buttonText: 'Contineus',
    buttinLink: "/Promotion",
    buttonVariant: 'contained',
  },

  {
    image: "https://scontent.fnak3-1.fna.fbcdn.net/v/t1.0-9/65875628_1897790406988617_2923525088097599488_o.jpg?_nc_cat=101&ccb=3&_nc_sid=174925&_nc_eui2=AeHKNVk15ZhXeXkHIT9E6fXH7EzLmcT18SPsTMuZxPXxI-JmpV4PjHPEIOFFv9y3Ss_AKab1r8JR79dKCXWGYZ9f&_nc_ohc=acejbMh5HtcAX83TqZy&_nc_ht=scontent.fnak3-1.fna&oh=97899499483f6784e679a5aecb39463f&oe=604DC96C",
    title: 'ระบบค้นหาบันทึกเงินเดือนพนักงาน',
    description: [/*'ระบบย่อย', 'บันทึกเงินเดือนพนักงาน', 'ตารางเวลาทำงานพนักงาน'*/],
    buttonText: 'Contineus',
    buttinLink: "/Salary/Search",
    buttonVariant: 'contained',
  },
  {
    image: "https://scontent.fnak3-1.fna.fbcdn.net/v/t1.0-9/137680050_1662022977332970_4873515341148094221_o.jpg?_nc_cat=103&ccb=3&_nc_sid=09cbfe&_nc_eui2=AeEzcInymktHLsCLHBMeFNrUFH2Um8uj5tsUfZSby6Pm27nnYPvMJ4y2I2zx62ilGkHAeFRnyj0_bHA1YeOGKtbf&_nc_ohc=vU_rAO3TE3oAX-y97dC&_nc_ht=scontent.fnak3-1.fna&oh=7865079140f9ff09127c73fcc7cbdf2b&oe=604FC6A8",
    title: 'ระบบค้นหารายการสั่งซื้อสินค้า',
    description: [],
    buttonText: 'Contineus',
    buttinLink: "/searchorderproduct",
    buttonVariant: 'contained',
  },
];


export default function Pricing() {
  const classes = useStyles();
  var ck = new Cookies()
  //var cookieEmail = ck.GetCookie()
  //var cookieID = ck.GetID()
  var cookieName = ck.GetName()

  return (
    <React.Fragment>
      <CssBaseline />
      <AppBar position="static" color="default" elevation={0} className={classes.appBar}>
        <Toolbar className={classes.toolbar}>
          <Typography variant="h3" color="inherit" noWrap className={classes.toolbarTitle}>
            ยินดีต้อนรับเข้าสู่ระบบผู้จัดการ
          </Typography>
          <nav>
            <Link variant="button" color="textPrimary" href="https://github.com/sut63/team13" className={classes.link}>
            source code
            </Link>
            
          </nav>
          <Button href="#" color="primary" variant="outlined" className={classes.link}
          component={RouterLink}
          to="/signinorderproduct">
            Logout
          </Button>
        </Toolbar>
      </AppBar>
      {/* Hero unit */}
      <Container maxWidth="sm" component="main" className={classes.heroContent}>
        <Typography component="h1" variant="h2" align="center" color="textPrimary" gutterBottom>
          คุณ {cookieName}
        </Typography>
        {/*<Typography variant="h5" align="center" color="textSecondary" component="p">
          Quickly build an effective pricing table for your potential customers with this layout.
          It&apos;s built with default Material-UI components with little customization.
  </Typography>*/}
      </Container>
      {/* End hero unit */}
      <Container maxWidth="md" component="main">
        <Grid container spacing={5} alignItems="flex-end">
            {tiers.map((tier) =>  (
              <Grid item key={tier} xs={12} sm={6} md={3}>
                <Card className={classes.card}>
                  <CardMedia
                    className={classes.cardMedia}
                    image={tier.image}
                    title="Image title"
                  />
                  <CardContent className={classes.card}>
                  <ul>
                    {tier.description.map((line) => (
                      <Typography component="li" variant="subtitle1" align="center" key={line}>
                          {tier.title}
                      </Typography>
                    ))}
                  </ul>
                    <Typography>
                    {tier.title}
                    </Typography>
                  </CardContent>
                  <CardActions>
                    <Button size="small" color="primary" to={tier.buttinLink} component={RouterLink}>
                      {tier.buttonText}
                    </Button>

                  </CardActions>
                </Card>

              </Grid>

            ))}
        </Grid>
      </Container>
      
      {/* Footer }
      <Container maxWidth="md" component="footer" className={classes.footer}>
        <Grid container spacing={4} justify="space-evenly">
          {footers.map((footer) => (
            <Grid item xs={6} sm={3} key={footer.title}>
              <Typography variant="h6" color="textPrimary" gutterBottom>
                {footer.title}
              </Typography>
              <ul>
                {footer.description.map((item) => (
                  <li key={item}>
                    <Link href="#" variant="subtitle1" color="textSecondary">
                      {item}
                    </Link>
                  </li>
                ))}
              </ul>
            </Grid>
          ))}
        </Grid>
        <Box mt={5}>
          <Copyright />
        </Box>
      </Container>
      {/* End footer */}
    </React.Fragment>
  );
}