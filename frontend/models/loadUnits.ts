import { Unit, UnitProcedure, Value, Parameter } from './units'
import datasBeding from './datas/Q2000/UP_Q2000_BEDING.json'
import datasDest from './datas/Q2000/UP_Q2000_DEST.json'
import datasEnde from './datas/Q2000/UP_Q2000_ENDE.json'
import datasEntl from './datas/Q2000/UP_Q2000_ENTL.json'

const upBeding: UnitProcedure = {
    name: datasBeding.name,
    time: datasBeding.time,
    author: datasBeding.author,
    desc: datasBeding.desc,
    params: datasBeding.params,
}
const upDest: UnitProcedure = {
    name: datasDest.name,
    time: datasDest.time,
    author: datasDest.author,
    desc: datasDest.desc,
    params: datasDest.params,
}
const upEnde: UnitProcedure = {
    name: datasEnde.name,
    time: datasEnde.time,
    author: datasEnde.author,
    desc: datasEnde.desc,
    params: datasEnde.params,
}
const upEntl: UnitProcedure = {
    name: datasEntl.name,
    time: datasEntl.time,
    author: datasEntl.author,
    desc: datasEntl.desc,
    params: datasEntl.params,
}
const ups: UnitProcedure[] = [upBeding, upDest, upEnde, upEntl]
export { ups }
