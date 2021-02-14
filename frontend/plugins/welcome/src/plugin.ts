import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignInOrderproduct from './components/orderproduct/SignInOrderproduct'
import Orderonline from './components/Orderonline'
import SignInOrderonline from './components/Orderonline/SignInOrderonline'
import Orderproduct from './components/orderproduct'
import Stock from './components/Stock'
import SplitsystemManager from './components/orderproduct/SplitsystemManager'
import Salary from './components/Salary'
import SalaryTable from './components/Salary/SalaryTable'
import Tablestock from './components/Stock/Tablestock'
import Tableorderproduct from './components/orderproduct/Tableorderproduct'
import LoginEmployee from './components/Stock/LoginEmployee'
import Promotion from './components/Promotion'
import Promotiontable from './components/Promotion/Promotiontable'
import SearchPromotion from './components/Promotion/SearchPromotion'
import EmployeeWorkingHours from './components/EmployeeWorkingHours'
import SearchEmployeeWorkingHours from './components/EmployeeWorkingHours/SearchEmployeeWorkingHours'
import TableEmployeeWorkingHours from './components/EmployeeWorkingHours/TableEmployeeWorkingHours'
import STable from './components/Salary/STable'
import SearchOrderproduct from './components/orderproduct/Searchorderproduct'
import Table from './components/Stock/Table'
import SearchOrderonline from './components/Orderonline/Searchorderonline'
import SearchSalary from './components/Salary/Search'
import Employeepage from './components/Stock/Employeepage'

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/SignInOrderproduct', SignInOrderproduct);
    router.registerRoute('/Orderonline', Orderonline);
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
    router.registerRoute('/SearchPromotion', SearchPromotion);
    router.registerRoute('/LoginEmployee', LoginEmployee);
    router.registerRoute('/EmployeeWorkingHours', EmployeeWorkingHours);
    router.registerRoute('/SearchEmployeeWorkingHours', SearchEmployeeWorkingHours);
    router.registerRoute('/TableEmployeeWorkingHours', TableEmployeeWorkingHours);
    router.registerRoute('/STable', STable);
    router.registerRoute('/Table', Table);
    router.registerRoute('/Salary/Search', SearchSalary);
    router.registerRoute('/Employeepage', Employeepage);
  },
});
