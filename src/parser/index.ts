import { readFhxFile } from "./readFile";

const fhx = {
    create: async (path: string) => await readFhxFile(path)
}

export default fhx;