import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntAssessment } from '../../api/models/EntAssessment';
import { EntEmployee } from '../../api/models/EntEmployee';
import { EntPosition } from '../../api/models/EntPosition';
import { EntSalary } from '../../api/models/EntSalary';
import { Grid } from '@material-ui/core';
import Typography from '@material-ui/core/Typography'
import AccountCircleIcon from '@material-ui/icons/AccountCircle';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);

/*const initialUserState = {
 name: 'System Analysis and Design',
 age: 20,
};
*/
interface salarys {
  Employee: string;
  Position: string;
  Assessment: string;
  Salarys: number;
  Bonus: number;
  SalaryTime: Date;
};

export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ระบบกรมธรรม์' };
  const api = new DefaultApi();
  //const [user, setUser] = useState(initialUserState);
  const [salarys, setSalary] = React.useState<Partial<EntSalary>>({});
  const [assessment, setAssessments] = React.useState<EntAssessment[]>([]);
  const [position, setPositions] = React.useState<EntPosition[]>([]);
  const [employee, setEmployees] = React.useState<EntEmployee[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [Assessment, setAssessmentid] = useState(Number);
  const [Position, setPositionid] = useState(Number);
  const [Emplayee, setEmplayeeid] = useState(Number);
  const [SalaryNumber, setSalarynumber] = useState(Number);
  const [Bonus, setBonus] = useState(Number);
  const [AddDate, setDatetime] = useState(String);

  useEffect(() => {

    const getEmployee = async () => {
      const res = await api.listEmployee({ limit: 10, offset: 0 });
      setLoading(false);
      setEmployees(res);
    };
    getEmployee();

    const getPosition = async () => {
      const res = await api.listPosition({ limit: 10, offset: 0 });
      setLoading(false);
      setPositions(res);
    };
    getPosition();

    const getAssessment = async () => {
      const res = await api.listAssessment({ limit: 10, offset: 0 });
      setLoading(false);
      setAssessments(res);
    };
    getAssessment();

  }, [loading]);

  const handleSalaryChange = (event: any) => {
    setSalarynumber(event.target.value as number);
  };

  const handleBonusChange = (event: any) => {
    setBonus(event.target.value as number);
  };

  const handleDatetimeChange = (event: any) => {
    setDatetime(event.target.value as string);
  };

  const CreateSalary = async () => {
    const Salary = {
      assessmentID: Assessment,
      employeeID: Emplayee,
      positionID: Position,
      bonus: Bonus,
      salarys: SalaryNumber,
      salaryDate: AddDate + ":00+07:00",
    }
    console.log(Salary);
    const res: any = await api.createSalary({ salary: Salary });
    setStatus(true)
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }

    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  const employee_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmplayeeid(event.target.value as number);
  };

  const assessment_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAssessmentid(event.target.value as number);
  };

  const position_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPositionid(event.target.value as number);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName || 'to Backstage'}`}
        subtitle=""
      >

        <Typography variant="h2" gutterBottom style={{ color: '#fafafa', marginLeft: 20 }}>
          <AccountCircleIcon />
        </Typography>

        <Typography variant="h4" gutterBottom style={{ color: '#fafafa', marginLeft: 40 }}>
        นายนนทกร พาอยู่สุข
      </Typography>

        <Link component={RouterLink} to="/">

          <Button variant="contained" color="secondary" style={{ marginLeft: 20 }}>

            Logout

           </Button>

        </Link>
      </Header>


      <Content>
        <ContentHeader title="สมัครกรมธรรม์">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกกรมธรรม์สำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    บันทึกกรมธรรม์ไม่สำเร็จ
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <table>
              <tr><td align="right">ชื่อพนักงาน</td><td>
                <FormControl
                  fullWidth
                  className={classes.margin}
                  variant="outlined"
                >
                  <InputLabel id="name_id-label">เลือกพนักงาน</InputLabel>
                  <Select
                    labelId="employee_id-label"
                    label="Employee"
                    id="employee_id"
                    value={Emplayee}
                    onChange={employee_id_handleChange}
                    style={{ width: 600 }}
                  >
                    {employee.map((item: EntEmployee) =>
                      <MenuItem value={item.id}>{item.name}</MenuItem>)}
                  </Select>
                </FormControl>
              </td></tr>


              <tr><td align="right">ตำแหน่ง</td><td>

                <FormControl
                  className={classes.margin}
                  variant="outlined"
                >
                  <InputLabel id="position_id-label">เลือกตำแหน่ง</InputLabel>
                  <Select
                    labelId="position_id-label"
                    label="position"
                    id="position_id"
                    value={Position}
                    onChange={position_id_handleChange}
                    style={{ width: 600 }}
                  >
                    {position.map((item: EntPosition) =>
                      <MenuItem value={item.id}>{item.position}</MenuItem>)}
                  </Select>
                </FormControl>
              </td></tr>

              <tr><td align="right">ผลการประเมิน</td><td>
                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >

                    <InputLabel id="assessments_id-label">เลือกผลการประเมิน</InputLabel>
                    <Select
                      labelId="assessment_id-label"
                      label="assessment"
                      id="assessment_id"
                      value={Assessment}
                      onChange={assessment_id_handleChange}
                      style={{ width: 600 }}
                    >
                      {assessment.map((item: EntAssessment) =>
                        <MenuItem value={item.id}>{item.assessment}</MenuItem>)}
                    </Select>
                  </FormControl>
                </div>
              </td></tr>

              <tr><td align="right">เงินเดือน</td><td>
                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >

                    <TextField id="outlined-number" type='number'  InputLabelProps={{
                  shrink: true,}}label="จำนวน" variant="outlined"
                  onChange = {handleSalaryChange}
                  />
                  </FormControl>
                </div>
              </td></tr>


              <tr><td align="right">โบนัส</td><td>
                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >

                    <TextField id="outlined-number" type='number'  InputLabelProps={{
                  shrink: true,}}label="F[oyl" variant="outlined"
                  onChange = {handleBonusChange}
                  />
                  </FormControl>

                </div>
              </td></tr>

              <tr><td align="right">วันที่เงินเดือนออก</td><td>
                <div>
                  <FormControl className={classes.margin} >
                    <TextField
                      id="AddDate"
                      label="AddDate"
                      type="datetime-local"
                      value={AddDate}
                      onChange={handleDatetimeChange}
                      //defaultValue="2017-05-24"
                      className={classes.textField}
                      InputLabelProps={{
                        shrink: true,
                      }}
                    />
                  </FormControl>

                </div>
              </td></tr>
            </table>

            <div className={classes.margin}>

              <Button
                onClick={() => {
                  CreateSalary();

                }}
                /* component={RouterLink}
                 to="/main"*/
                variant="contained"

                color="primary"
                style={{ marginLeft : 270}}
              >
                บันทึกเงินเดือน
                                </Button>
              <Link component={RouterLink} to="/WelcomePage">

                <Button variant="contained" color="secondary" style={{ marginLeft : 40}}>

                  หน้าหลัก

           </Button>

              </Link>

            </div>

          </form>
        </div>
      </Content>
    </Page>

  );
}