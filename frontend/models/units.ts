interface Unit {
    desc: string
    unitname: string
    unitprocedure: UnitProcedure[]
}

interface UnitProcedure {
    name: string
    time: number
    author: string
    desc: string
    params: Parameter[]
}

interface Parameter {
    name: string
    desc: string
    value: Value
}

interface Value {
    stringvalue: string
    set: string
    high: string
    low: string
    cv: string
    unit: string
}

export { Unit, UnitProcedure, Parameter, Value }