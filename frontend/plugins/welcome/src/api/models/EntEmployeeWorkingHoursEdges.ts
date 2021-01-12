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
import {
    EntDay,
    EntDayFromJSON,
    EntDayFromJSONTyped,
    EntDayToJSON,
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
    EntRole,
    EntRoleFromJSON,
    EntRoleFromJSONTyped,
    EntRoleToJSON,
    EntShift,
    EntShiftFromJSON,
    EntShiftFromJSONTyped,
    EntShiftToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntEmployeeWorkingHoursEdges
 */
export interface EntEmployeeWorkingHoursEdges {
    /**
     * 
     * @type {EntDay}
     * @memberof EntEmployeeWorkingHoursEdges
     */
    day?: EntDay;
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntEmployeeWorkingHoursEdges
     */
    employeeWorkingHours?: EntEmployee;
    /**
     * 
     * @type {EntRole}
     * @memberof EntEmployeeWorkingHoursEdges
     */
    role?: EntRole;
    /**
     * 
     * @type {EntShift}
     * @memberof EntEmployeeWorkingHoursEdges
     */
    shift?: EntShift;
}

export function EntEmployeeWorkingHoursEdgesFromJSON(json: any): EntEmployeeWorkingHoursEdges {
    return EntEmployeeWorkingHoursEdgesFromJSONTyped(json, false);
}

export function EntEmployeeWorkingHoursEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntEmployeeWorkingHoursEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'day': !exists(json, 'day') ? undefined : EntDayFromJSON(json['day']),
        'employeeWorkingHours': !exists(json, 'employeeWorkingHours') ? undefined : EntEmployeeFromJSON(json['employeeWorkingHours']),
        'role': !exists(json, 'role') ? undefined : EntRoleFromJSON(json['role']),
        'shift': !exists(json, 'shift') ? undefined : EntShiftFromJSON(json['shift']),
    };
}

export function EntEmployeeWorkingHoursEdgesToJSON(value?: EntEmployeeWorkingHoursEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'day': EntDayToJSON(value.day),
        'employeeWorkingHours': EntEmployeeToJSON(value.employeeWorkingHours),
        'role': EntRoleToJSON(value.role),
        'shift': EntShiftToJSON(value.shift),
    };
}

