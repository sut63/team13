import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo'
import SignIn from './components/SignIn'
import SignInOrderproduct from './components/orderproduct/SignInOrderproduct'
import Orderonline from './components/Orderonline'
import Orderonlinetable from './components/Orderonline/Tableorderonline'
import SignInOrderonline from './components/Orderonline/SignInOrderonline'
import Orderproduct from './components/orderproduct'
import Stock from './components/Stock'
import SplitsystemManager from './components/SplitsystemManager'
import Salary from './components/Salary'
import SalaryTable from './components/SalaryTable'
import Tablestock from './components/Tablestock'
import Tableorderproduct from './components/orderproduct/Tableorderproduct'
import LoginEmployee from './components/Stock/LoginEmployee'
import ManagerPage from './components/ManagerPage'
import Promotion from './components/Promotion'
import Promotiontable from './components/Promotion/Promotiontable'
import EmployeeWorkingHours from './components/EmployeeWorkingHours'
import SearchEmployeeWorkingHours from './components/SearchEmployeeWorkingHours'
import TableEmployeeWorkingHours from './components/TableEmployeeWorkingHours'
import STable from './components/STable'
import SearchOrderproduct from './components/orderproduct/Searchorderproduct'
import Table from './components/Stock/Table'
import SearchOrderonline from './components/Orderonline/Searchorderonline'
import SearchSalary from './components/Salary/Search'

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/SignInOrderproduct', SignInOrderproduct);
    router.registerRoute('/Orderonline', Orderonline);
    router.registerRoute('/Orderonlinetable', Orderonlinetable);
    router.registerRoute('/SignInOrderonline', SignInOrderonline);
    router.registerRoute('/SearchOrderonline', SearchOrderonline);
    router.registerRoute('/Orderproduct', Orderproduct);
    router.registerRoute('/SearchOrderproduct', SearchOrderproduct);
    router.registerRoute('/Stock', Stock);
    router.registerRoute('/SplitsystemManager', SplitsystemManager);
    router.registerRoute('/Salary', Salary);
    router.registerRoute('/SalaryTable', SalaryTable);
    router.registerRoute('/Tablestock', Tablestock);
    router.registerRoute('/Tableorderproduct', Tableorderproduct);
    router.registerRoute('/Promotion', Promotion);
    router.registerRoute('/Promotiontable', Promotiontable);
    router.registerRoute('/LoginEmployee', LoginEmployee);
    router.registerRoute('/ManagerPage', ManagerPage);
    router.registerRoute('/EmployeeWorkingHours', EmployeeWorkingHours);
    router.registerRoute('/SearchEmployeeWorkingHours', SearchEmployeeWorkingHours);
    router.registerRoute('/TableEmployeeWorkingHours', TableEmployeeWorkingHours);
    router.registerRoute('/STable', STable);
    router.registerRoute('/Table', Table);
    router.registerRoute('/Salary/Search', SearchSalary);
  },
});
