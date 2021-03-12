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
 * @interface ControllersEmployeeWorkingHours
 */
export interface ControllersEmployeeWorkingHours {
    /**
     * 
     * @type {number}
     * @memberof ControllersEmployeeWorkingHours
     */
    beginWork?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersEmployeeWorkingHours
     */
    codeWork?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersEmployeeWorkingHours
     */
    day?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersEmployeeWorkingHours
     */
    employee?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersEmployeeWorkingHours
     */
    getOffWork?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersEmployeeWorkingHours
     */
    idemployee?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersEmployeeWorkingHours
     */
    idnumber?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersEmployeeWorkingHours
     */
    role?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersEmployeeWorkingHours
     */
    wages?: number;
}

export function ControllersEmployeeWorkingHoursFromJSON(json: any): ControllersEmployeeWorkingHours {
    return ControllersEmployeeWorkingHoursFromJSONTyped(json, false);
}

export function ControllersEmployeeWorkingHoursFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersEmployeeWorkingHours {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'beginWork': !exists(json, 'beginWork') ? undefined : json['beginWork'],
        'codeWork': !exists(json, 'codeWork') ? undefined : json['codeWork'],
        'day': !exists(json, 'day') ? undefined : json['day'],
        'employee': !exists(json, 'employee') ? undefined : json['employee'],
        'getOffWork': !exists(json, 'getOffWork') ? undefined : json['getOffWork'],
        'idemployee': !exists(json, 'idemployee') ? undefined : json['idemployee'],
        'idnumber': !exists(json, 'idnumber') ? undefined : json['idnumber'],
        'role': !exists(json, 'role') ? undefined : json['role'],
        'wages': !exists(json, 'wages') ? undefined : json['wages'],
    };
}

export function ControllersEmployeeWorkingHoursToJSON(value?: ControllersEmployeeWorkingHours | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'beginWork': value.beginWork,
        'codeWork': value.codeWork,
        'day': value.day,
        'employee': value.employee,
        'getOffWork': value.getOffWork,
        'idemployee': value.idemployee,
        'idnumber': value.idnumber,
        'role': value.role,
        'wages': value.wages,
    };
}


