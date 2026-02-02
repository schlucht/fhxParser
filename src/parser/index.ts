import { readFhxFile } from "./readFile";

const fhx = {
    create: (path: string) => readFhxFile(path)
}

export default fhx;