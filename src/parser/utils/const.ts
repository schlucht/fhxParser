
export const Regex = {
    user: /user="([^"]+)"/,
    time: /time=([0-9]+)/,
    timeComment: /\/\*\s*"([^"]+)"\s*\*\//,

    version: /VERSION=([0-9]+)(?:\/\*\s*"([^"]+)"\s*\*\/)?/,
    major: /MAJOR_VERSION=([0-9]+)/,
    minor: /MINOR_VERSION=([0-9]+)/,
    maintenance: /MAINTENANCE_VERSION=([0-9]+)/,
    build: /BUILD_VERSION=([0-9]+)/,
    buildId: /BUILD_ID="([^"]+)"/,
    versionStr: /VERSION_STR="([^"]+)"/,
    upgrade: /ONLINE_UPGRADE=([A-Z])/,
    local: /LOCALE="([^"]+)"/,
    meta: /user="([^"]+)"\s+time=([0-9]+)(?:\/\*\s*"([^"]+)"\s*\*\/)?/,

    formParam: /FORMULA_PARAMETER NAME="([^"]+)"\s+TYPE=([^"]+)/,
    connection: /CONNECTION=([^"]+)/,
    group: /GROUP="([^"]+)/,

    stringAttr: {
        set: /SET="([^"]+)/,
        sv: /STRING_VALUE="([^"]+)/,
        ch: /CHANGEABLE=([A-Z])/,
    },

    numAttr: {
       numeric: /DESCRIPTION="([^"]*)"\s+HIGH=(-?\d+)\s+LOW=(-?\d+)\s+SCALABLE=([A-Z])\s+CV=(\d+)\s+UNITS="([^"]+)"/
    },
    
    cv: /CV="([^"]+)/,

    name: /NAME="([^"]+)/,
    type: /TYPE=([\w]+)/,
    category: /CATEGORY="([^"]+)/,
    description: /DESCRIPTION="([^"]+)"/,
    useEquipmentTrains: /USE_EQUIPMENT_TRAINS=([A-Z])/,
    equipmentUnitClass: /EQUIPMENT_UNIT_CLASS="([^"]+)/,
    author: /AUTHOR="([^"]+)/,
    abstract: /ABSTRACT="([^"]+)/,
    batchUnits: /BATCH_UNITS="([^"]+)/,
    batchLength: /BATCH_LENGTH="([^"]+)/,
    defaultBatchSize: /DEFAULT_BATCH_SIZE=([0-9]+)/,
    minimumBatchSize: /MINIMUM_BATCH_SIZE=([0-9]+)/,
    maximumBatchSize: /MAXIMUM_BATCH_SIZE=([0-9]+)/,
    productCode: /PRODUCT_CODE="([^"]+)/,
    productName: /PRODUCT_NAME="([^"]+)/,
    recipeApproval: /RECIPE_APPROVAL_INFO="([^"]+)/,
    versionRecipe: /VERSION="([^"]+)/,

} as const;

export const Blocks = {
    version: 'VERSION',
    schema: 'SCHEMA',
    local: 'LOCALE',
    recipe: 'BATCH_RECIPE',
} as const;