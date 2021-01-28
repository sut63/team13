/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersSalary
 */
export interface ControllersSalary {
    /**
     * 
     * @type {string}
     * @memberof ControllersSalary
     */
    accountNumber?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersSalary
     */
    assessmentID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersSalary
     */
    bonus?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersSalary
     */
    employeeID?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersSalary
     */
    idemployee?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersSalary
     */
    positionID?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersSalary
     */
    salaryDate?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersSalary
     */
    salarys?: number;
}

export function ControllersSalaryFromJSON(json: any): ControllersSalary {
    return ControllersSalaryFromJSONTyped(json, false);
}

export function ControllersSalaryFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersSalary {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'accountNumber': !exists(json, 'accountNumber') ? undefined : json['accountNumber'],
        'assessmentID': !exists(json, 'assessmentID') ? undefined : json['assessmentID'],
        'bonus': !exists(json, 'bonus') ? undefined : json['bonus'],
        'employeeID': !exists(json, 'employeeID') ? undefined : json['employeeID'],
        'idemployee': !exists(json, 'idemployee') ? undefined : json['idemployee'],
        'positionID': !exists(json, 'positionID') ? undefined : json['positionID'],
        'salaryDate': !exists(json, 'salaryDate') ? undefined : json['salaryDate'],
        'salarys': !exists(json, 'salarys') ? undefined : json['salarys'],
    };
}

export function ControllersSalaryToJSON(value?: ControllersSalary | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'accountNumber': value.accountNumber,
        'assessmentID': value.assessmentID,
        'bonus': value.bonus,
        'employeeID': value.employeeID,
        'idemployee': value.idemployee,
        'positionID': value.positionID,
        'salaryDate': value.salaryDate,
        'salarys': value.salarys,
    };
}


