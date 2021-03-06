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
    EntEndWork,
    EntEndWorkFromJSON,
    EntEndWorkFromJSONTyped,
    EntEndWorkToJSON,
    EntRole,
    EntRoleFromJSON,
    EntRoleFromJSONTyped,
    EntRoleToJSON,
    EntStartWork,
    EntStartWorkFromJSON,
    EntStartWorkFromJSONTyped,
    EntStartWorkToJSON,
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
    employee?: EntEmployee;
    /**
     * 
     * @type {EntEndWork}
     * @memberof EntEmployeeWorkingHoursEdges
     */
    endwork?: EntEndWork;
    /**
     * 
     * @type {EntRole}
     * @memberof EntEmployeeWorkingHoursEdges
     */
    role?: EntRole;
    /**
     * 
     * @type {EntStartWork}
     * @memberof EntEmployeeWorkingHoursEdges
     */
    startwork?: EntStartWork;
}

export function EntEmployeeWorkingHoursEdgesFromJSON(json: any): EntEmployeeWorkingHoursEdges {
    return EntEmployeeWorkingHoursEdgesFromJSONTyped(json, false);
}

export function EntEmployeeWorkingHoursEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntEmployeeWorkingHoursEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'day': !exists(json, 'Day') ? undefined : EntDayFromJSON(json['Day']),
        'employee': !exists(json, 'Employee') ? undefined : EntEmployeeFromJSON(json['Employee']),
        'endwork': !exists(json, 'Endwork') ? undefined : EntEndWorkFromJSON(json['Endwork']),
        'role': !exists(json, 'Role') ? undefined : EntRoleFromJSON(json['Role']),
        'startwork': !exists(json, 'Startwork') ? undefined : EntStartWorkFromJSON(json['Startwork']),
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
        'employee': EntEmployeeToJSON(value.employee),
        'endwork': EntEndWorkToJSON(value.endwork),
        'role': EntRoleToJSON(value.role),
        'startwork': EntStartWorkToJSON(value.startwork),
    };
}


