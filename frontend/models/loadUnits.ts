import { Unit, UnitProcedure, Value, Parameter } from './units'
import datas from './units.json'

const Units: Unit[] = []

datas.forEach( (data: Unit) => {
    Units.push(data)
})


export { Units }