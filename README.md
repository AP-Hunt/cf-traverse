# cf-traverse

`cf-traverse` is a [Cloud Foundry CLI](https://github.com/cloudfoundry/cli) plugin for traversing the relationships within
the Cloud Foundry API. For instance, getting the service offering from a given service instance name or GUID.

It was born out of the frustration of writing code for the same handful of API traversals every time they were needed. The
goal is to make it easier to write scripts which work with the Cloud Foundry API to get data out.

## Usage

```
cf traverse [SOURCE ENTITY TYPE] [RELATION] [SOURCE ENTITY IDENTIFIER] 
```

Traverse takes 3 arguments: the type of the source entity (e.g. service instance), the relation to find, and an identifier 
for the source entity. The table below shows all the supported source entities, relations, and what qualifies as an identifier

<table>
    <thead>
        <tr>
            <th>Source entity type</th>
            <th>Source entity identifier</th>
            <th>Relation</th>
        </tr>    
    </thead>
    <tbody>
        <!-- Service instance -->
        <tr>
            <td rowspan="5">service_instance</td>
            <td rowspan="5">A service instance guid, or service instance name</td>
        </tr>
        <tr><td>space</td></tr>
        <tr><td>org</td></tr>
        <tr><td>plan</td></tr>
        <tr><td>service_offering</td></tr>
        <!-- Service plan -->
        <tr>
            <td rowspan="2">service_plan</td>
            <td rowspan="2">A service plan guid</td>
        </tr>
        <tr><td>instances_of</td></tr>
        <!-- Service offering -->
        <tr>
            <td rowspan="2">service_offering</td>
            <td rowspan="2">A service offering guid, or service offering name</td>
        </tr>
        <tr><td>instances_of</td></tr>
    </tbody>
</table>

## Installation
The current best path is to run `make install`. Releases in GitHub will come in the future.

