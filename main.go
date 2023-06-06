package main

import (
	"fmt"

	"github.com/schlucht/fhxreader/fhx-app/server"
)

var txt = `/* Version: 14.3.1.7332.xr */
/* "06-Dec-2022 08:06:24" */

SCHEMA
 user="DLEUTHARDT" time=1670310224/* "06-Dec-2022 08:03:44" */
{
  /* Database last updated on "30-Jul-2019 03:07:54" */
  VERSION=1564448874/* "30-Jul-2019 03:07:54" */
  MAJOR_VERSION=14
  MINOR_VERSION=3
  MAINTENANCE_VERSION=1
  BUILD_VERSION=7332
  BUILD_ID="xr"
  VERSION_STR="14.3.1.7332.xr"
  ONLINE_UPGRADE=F
}
LOCALE
 user="DLEUTHARDT" time=1670310224/* "06-Dec-2022 08:03:44" */
{
  LOCALE="English_United States.1252"
}
BATCH_RECIPE NAME="UP_Q2800_START" TYPE=UNIT_PROCEDURE CATEGORY="Recipes/Unit Procedures/STANDARD_Q2800"
 user="LSCHMID1" time=1665401910/* "10-Oct-2022 13:38:30" */
{
  DESCRIPTION="Q2800 Start"
  USE_EQUIPMENT_TRAINS=F
  EQUIPMENT_UNIT="Q2800"
  ENFORCED_FORMULA_SELECTION_ENABLED=F
  DEFAULT_SELECTED_FORMULA=""
  PARAMETERS_LOCKED_BY_DEFAULT=F
  AUTHOR="Schmid Lothar"
  ABSTRACT=""
  BATCH_UNITS=""
  BATCH_LENGTH=""
  DEFAULT_BATCH_SIZE=100
  MINIMUM_BATCH_SIZE=1
  MAXIMUM_BATCH_SIZE=100
  PRODUCT_CODE=""
  PRODUCT_NAME="Allegmeine Standart UP"
  RECIPE_APPROVAL_INFO=""
  VERSION="2022.3"
  FORMULA_PARAMETER NAME="FP_HK_OPT" TYPE=ENUMERATION_VALUE
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    DESCRIPTION="Option Mantel oder Innentemp"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_HK_OPT_HO" TYPE=ENUMERATION_VALUE
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    DESCRIPTION="Option Mantel oder Innentemp bei SS"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_HK_TCW" TYPE=BATCH_PARAMETER_REAL
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    DESCRIPTION="Sollwert HKK für die HK Option"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_HK_TCW_HO" TYPE=BATCH_PARAMETER_REAL
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    DESCRIPTION="Sollwert HKK für die HK Option  bei SS"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_KK_TCW" TYPE=BATCH_PARAMETER_REAL
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    GROUP="Operating"
    DESCRIPTION="Temp. Sollwert Kühler"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_KK_OPT" TYPE=ENUMERATION_VALUE
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    GROUP="Operating"
    DESCRIPTION="Ein- Ausschalten KK Kreislauf"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_KK_TAH" TYPE=BATCH_PARAMETER_REAL
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    GROUP="Operating"
    DESCRIPTION="Temp. Alarm H Kühler"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_KK_TAHH" TYPE=BATCH_PARAMETER_REAL
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    GROUP="Operating"
    DESCRIPTION="Temp. Alarm HH Kühler"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_APP_SOLE" TYPE=BATCH_PARAMETER_REAL
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    DESCRIPTION="Umschaltpunkt kalter zu warmer Sole vom HKK"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_ABGAS_TAH" TYPE=BATCH_PARAMETER_REAL
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    GROUP="Operating"
    DESCRIPTION="Austrittstemp. Abgas H"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_ABGAS_TAHH" TYPE=BATCH_PARAMETER_REAL
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    GROUP="Operating"
    DESCRIPTION="Austrittstemp. Abgas HH Alarm löst SS aus"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_SIS_EIN" TYPE=ENUMERATION_VALUE
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    GROUP="Operating"
    DESCRIPTION="TRUE kontrolliert ob richtiger SIS Parameter eingestellt ist."
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_SIS_PRODUKT" TYPE=UNICODE_STRING
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    DESCRIPTION="SIS Parameter Produkt  (Wenn kein SIS leer lassen)"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_APP_ABLUFT_1" TYPE=ENUMERATION_VALUE
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    DESCRIPTION="Wäscher bei Störung vom Wäscher geht RW in SS"
    IS_PARAMETER_LOCKED=F
  }
  FORMULA_PARAMETER NAME="FP_APP_ABLUFT_2" TYPE=ENUMERATION_VALUE
  {
    CONNECTION=INPUT
    RECTANGLE= { X=-50 Y=-50 H=1 W=1 }
    DESCRIPTION="Wäscher bei Störung vom Wäscher geht RW in SS"
    IS_PARAMETER_LOCKED=F
  }
  ATTRIBUTE_INSTANCE NAME="FP_HK_OPT"
  {
    VALUE
    {
      SET="LGF_HK_OPT"
      STRING_VALUE="MANTELTEMP"
      CHANGEABLE=F
    }
  }
  ATTRIBUTE_INSTANCE NAME="FP_HK_OPT_HO"
  {
    VALUE
    {
      SET="LGF_HK_OPT"
      STRING_VALUE="MANTELTEMP"
      CHANGEABLE=F
    }
  }
  ATTRIBUTE_INSTANCE NAME="FP_HK_TCW"
  {
    VALUE { DESCRIPTION="" HIGH=180 LOW=-20 SCALABLE=F CV=25 UNITS="°C" }
  }
  ATTRIBUTE_INSTANCE NAME="FP_HK_TCW_HO"
  {
    VALUE { DESCRIPTION="" HIGH=180 LOW=-20 SCALABLE=F CV=25 UNITS="°C" }
  }
  ATTRIBUTE_INSTANCE NAME="FP_KK_TCW"
  {
    VALUE { DESCRIPTION="" HIGH=100 LOW=-20 SCALABLE=F CV=35 UNITS="°C" }
  }
  ATTRIBUTE_INSTANCE NAME="FP_KK_OPT"
  {
    VALUE
    {
      SET="L_EIN_AUS"
      STRING_VALUE="EIN"
      CHANGEABLE=F
    }
  }
  ATTRIBUTE_INSTANCE NAME="FP_KK_TAH"
  {
    VALUE { DESCRIPTION="" HIGH=100 LOW=-20 SCALABLE=F CV=40 UNITS="°C" }
  }
  ATTRIBUTE_INSTANCE NAME="FP_KK_TAHH"
  {
    VALUE { DESCRIPTION="" HIGH=100 LOW=-20 SCALABLE=F CV=45 UNITS="°C" }
  }
  ATTRIBUTE_INSTANCE NAME="FP_APP_SOLE"
  {
    VALUE { DESCRIPTION="" HIGH=50 LOW=-20 SCALABLE=F CV=45 UNITS="°C" }
  }
  ATTRIBUTE_INSTANCE NAME="FP_ABGAS_TAH"
  {
    VALUE { DESCRIPTION="" HIGH=100 LOW=0 SCALABLE=F CV=80 UNITS="°C" }
  }
  ATTRIBUTE_INSTANCE NAME="FP_ABGAS_TAHH"
  {
    VALUE { DESCRIPTION="" HIGH=100 LOW=0 SCALABLE=F CV=85 UNITS="°C" }
  }
  ATTRIBUTE_INSTANCE NAME="FP_SIS_EIN"
  {
    VALUE
    {
      SET="LGF_BEDING_DI"
      STRING_VALUE="False"
      CHANGEABLE=F
    }
  }
  ATTRIBUTE_INSTANCE NAME="FP_SIS_PRODUKT"
  {
    VALUE { CV="Q2800_SULF_SIS" }
  }
  ATTRIBUTE_INSTANCE NAME="FP_APP_ABLUFT_1"
  {
    VALUE
    {
      SET="LP_ABLUFT"
      STRING_VALUE="kein Abluftsystem"
      CHANGEABLE=F
    }
  }
  ATTRIBUTE_INSTANCE NAME="FP_APP_ABLUFT_2"
  {
    VALUE
    {
      SET="LP_ABLUFT"
      STRING_VALUE="kein Abluftsystem"
      CHANGEABLE=F
    }
  }
  PFC_ALGORITHM
  {
    STEP NAME="OP_BEDING_DI:1" DEFINITION="OP_BEDING_DI"
    {
      DESCRIPTION="SIS Abfragen"
      RECTANGLE= { X=260 Y=150 H=40 W=180 }
      STEP_PARAMETER NAME="FP_BESCHREI_TEXT"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_FSB_MESSTELLE"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_SIS_PRODUKT"
      }
      STEP_PARAMETER NAME="FP_FSB_VERGLEICH"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_SIS_EIN"
        GROUP="Operating"
      }
      KEY_PARAMETER=""
    }
    STEP NAME="OP_BEGAS:1" DEFINITION="OP_BEGAS"
    {
      DESCRIPTION="Normal"
      RECTANGLE= { X=50 Y=390 H=40 W=180 }
      STEP_PARAMETER NAME="FP_BESCHREI_TEXT"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_FSB_PSH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_FSB_PSL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_OPTION"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_OPTION_HO"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_OPTION_PC"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_RAMPE"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_YS_AUF_PSH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_YS_AUF_PSH_HO"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_YS_AUF_PSL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_YS_AUF_PSL_HO"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      ATTRIBUTE_INSTANCE NAME="FP_BESCHREI_TEXT"
      {
        VALUE { CV="Normal" }
      }
      KEY_PARAMETER="FP_YS_AUF_PSH"
    }
    STEP NAME="OP_DRUCK:1" DEFINITION="OP_DRUCK"
    {
      DESCRIPTION="Normal"
      RECTANGLE= { X=50 Y=290 H=40 W=180 }
      STEP_PARAMETER NAME="FP_BESCHREI_TEXT"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_FSB_PSH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_FSB_PSL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_OPTION"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_OPTION_VAK_HO"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_PAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_PAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_PAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_PALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_PC_W"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_PC_W_HO"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_PV_KENNLINIE"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_RAMPE"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      ATTRIBUTE_INSTANCE NAME="FP_RAMPE"
      {
        VALUE { CV=0 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_BESCHREI_TEXT"
      {
        VALUE { CV="Normal" }
      }
      KEY_PARAMETER="FP_PC_W"
    }
    STEP NAME="OP_HK:1" DEFINITION="OP_HK"
    {
      DESCRIPTION="Mantel 30°C"
      RECTANGLE= { X=480 Y=290 H=40 W=180 }
      STEP_PARAMETER NAME="FP_BESCHREI_TEXT"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_FSB_TSH"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_FSB_TSL"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_MANENDE"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_OPTION"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_HK_OPT"
      }
      STEP_PARAMETER NAME="FP_OPTION_HO"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_HK_OPT_HO"
      }
      STEP_PARAMETER NAME="FP_RAMPE"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_TAH_BRUEDEN"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_TAH_HKK"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_TAH_INNEN"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_TAHH_HKK"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_TAHH_INNEN"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_TAL_HKK"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_TAL_INNEN"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_TALL_HKK"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_TALL_INNEN"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_TC_W"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_HK_TCW"
      }
      STEP_PARAMETER NAME="FP_TC_W_HO"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_HK_TCW_HO"
      }
      STEP_PARAMETER NAME="FP_TT_INNEN"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_TT_INNEN_HO"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_W_BGR_D_H"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_W_BGR_D_L"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_W_BGR_D_L_HO"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_W_BGR_H"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_W_BGR_L"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_W_BGR_L_HO"
      {
        ORIGIN=CONSTANT
      }
      ATTRIBUTE_INSTANCE NAME="FP_TAH_INNEN"
      {
        VALUE { CV=110 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_TAHH_INNEN"
      {
        VALUE { CV=120 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_TC_W"
      {
        VALUE { CV=30 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_TC_W_HO"
      {
        VALUE { CV=30 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_W_BGR_D_H"
      {
        VALUE { CV=60 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_W_BGR_D_L"
      {
        VALUE { CV=40 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_W_BGR_D_L_HO"
      {
        VALUE { CV=60 }
      }
      KEY_PARAMETER="FP_TC_W"
    }
    STEP NAME="OP_KONDENS:1" DEFINITION="OP_KONDENS"
    {
      DESCRIPTION="30°C"
      RECTANGLE= { X=480 Y=390 H=40 W=180 }
      STEP_PARAMETER NAME="FP_BESCHREI_TEXT"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_OPTION"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_KK_OPT"
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_TAH_AUSTRITT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_TAH_KK"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_KK_TAH"
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_TAHH_AUSTRITT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_TAHH_KK"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_KK_TAHH"
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_TAL_KK"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_TC_W"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_KK_TCW"
        GROUP="Operating"
      }
      ATTRIBUTE_INSTANCE NAME="FP_TAH_KK"
      {
        VALUE { CV=35 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_TAHH_KK"
      {
        VALUE { CV=40 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_TC_W"
      {
        VALUE { CV=30 }
      }
      KEY_PARAMETER="FP_TC_W"
    }
    STEP NAME="OP_MELDEN:1" DEFINITION="OP_MELDEN"
    {
      DESCRIPTION="Startkontrollen"
      RECTANGLE= { X=50 Y=150 H=40 W=180 }
      STEP_PARAMETER NAME="FP_BESCHREI_TEXT"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_MELDETEXT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      ATTRIBUTE_INSTANCE NAME="FP_MELDETEXT"
      {
        VALUE { CV="Kontrollen beendet?" }
      }
      ATTRIBUTE_INSTANCE NAME="FP_BESCHREI_TEXT"
      {
        VALUE { CV="Startkontrollen" }
      }
      KEY_PARAMETER="FP_BESCHREI_TEXT"
    }
    STEP NAME="OP_SET_AKTR_25:1" DEFINITION="OP_SET_AKTR_25"
    {
      DESCRIPTION="Bodenventil schliessen"
      RECTANGLE= { X=260 Y=290 H=40 W=180 }
      STEP_PARAMETER NAME="FP_AK01"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK01_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK01_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK02"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK02_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK02_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK03"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK03_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK03_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK04"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK04_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK04_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK05"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK05_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK05_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK06"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK06_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK06_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK07"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK07_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK07_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK08"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK08_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK08_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK09"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK09_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK09_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK10"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK10_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK10_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK11"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK11_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK11_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK12"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK12_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK12_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK13"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK13_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK13_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK14"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK14_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK14_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK15"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK15_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK15_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK16"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK16_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK16_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK17"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK17_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK17_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK18"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK18_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK18_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK19"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK19_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK19_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK20"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK20_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK20_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK21"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK21_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK21_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK22"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK22_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK22_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK23"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK23_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK23_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK24"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK24_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK24_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_AK25"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_AK25_OPT_AKT"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_AK25_OPT_MODE"
      {
        ORIGIN=CONSTANT
        GROUP="hidden"
      }
      STEP_PARAMETER NAME="FP_BESCHREI_TEXT"
      {
        ORIGIN=CONSTANT
      }
      ATTRIBUTE_INSTANCE NAME="FP_AK01"
      {
        VALUE { CV="Q2800YS6020" }
      }
      ATTRIBUTE_INSTANCE NAME="FP_AK01_OPT_AKT"
      {
        VALUE
        {
          SET="LGF_AKTO_OPT_AKT"
          STRING_VALUE="VENTIL ZU"
          CHANGEABLE=F
        }
      }
      ATTRIBUTE_INSTANCE NAME="FP_AK02"
      {
        VALUE { CV="Q2800YS6000" }
      }
      ATTRIBUTE_INSTANCE NAME="FP_AK02_OPT_AKT"
      {
        VALUE
        {
          SET="LGF_AKTO_OPT_AKT"
          STRING_VALUE="VENTIL ZU"
          CHANGEABLE=F
        }
      }
      ATTRIBUTE_INSTANCE NAME="FP_BESCHREI_TEXT"
      {
        VALUE { CV="Falls offen schliessen" }
      }
      KEY_PARAMETER=""
    }
    STEP NAME="OP_SET_ALAR_15:1" DEFINITION="OP_SET_ALAR_15"
    {
      DESCRIPTION="Alarme setzen"
      RECTANGLE= { X=260 Y=390 H=40 W=180 }
      STEP_PARAMETER NAME="FP_BESCHREI_TEXT"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_MST01"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST01_XAH"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_ABGAS_TAH"
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST01_XAHH"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_ABGAS_TAHH"
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST01_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST01_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST02"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST02_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST02_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST02_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST02_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST03"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST03_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST03_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST03_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST03_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST04"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST04_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST04_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST04_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST04_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST05"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST05_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST05_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST05_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST05_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST06"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST06_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST06_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST06_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST06_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST07"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST07_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST07_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST07_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST07_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST08"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST08_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST08_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST08_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST08_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST09"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST09_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST09_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST09_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST09_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST10"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST10_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST10_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST10_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST10_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST11"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST11_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST11_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST11_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST11_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST12"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST12_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST12_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST12_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST12_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST13"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST13_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST13_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST13_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST13_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST14"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST14_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST14_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST14_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST14_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST15"
      {
        ORIGIN=CONSTANT
        GROUP="Configuration"
      }
      STEP_PARAMETER NAME="FP_MST15_XAH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST15_XAHH"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST15_XAL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      STEP_PARAMETER NAME="FP_MST15_XALL"
      {
        ORIGIN=CONSTANT
        GROUP="Operating"
      }
      ATTRIBUTE_INSTANCE NAME="FP_MST01"
      {
        VALUE { CV="Q2800TI2530" }
      }
      ATTRIBUTE_INSTANCE NAME="FP_MST01_XAH"
      {
        VALUE { CV=80 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_MST01_XAHH"
      {
        VALUE { CV=85 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_MST02"
      {
        VALUE { CV="Q2800LI0010" }
      }
      ATTRIBUTE_INSTANCE NAME="FP_MST02_XAH"
      {
        VALUE { CV=11000 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_MST02_XAHH"
      {
        VALUE { CV=11000 }
      }
      KEY_PARAMETER=""
    }
    STEP NAME="OP_SET_AP_B:1" DEFINITION="OP_SET_AP_B"
    {
      DESCRIPTION="Normal"
      RECTANGLE= { X=50 Y=490 H=40 W=180 }
      STEP_PARAMETER NAME="FP_ABLUFT_1"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_APP_ABLUFT_1"
      }
      STEP_PARAMETER NAME="FP_ABLUFT_2"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_APP_ABLUFT_2"
      }
      STEP_PARAMETER NAME="FP_BESCHREI_TEXT"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_HK_TS_SOLE"
      {
        ORIGIN=DEFERRED
        DEFERRED_TO="FP_APP_SOLE"
      }
      STEP_PARAMETER NAME="FP_LAHH"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_LSHHH"
      {
        ORIGIN=CONSTANT
      }
      STEP_PARAMETER NAME="FP_VP"
      {
        ORIGIN=CONSTANT
      }
      ATTRIBUTE_INSTANCE NAME="FP_LAHH"
      {
        VALUE { CV=12000 }
      }
      ATTRIBUTE_INSTANCE NAME="FP_LSHHH"
      {
        VALUE { CV=12000 }
      }
      KEY_PARAMETER="FP_HK_TS_SOLE"
    }
    STEP NAME="START" DEFINITION=""
    {
      DESCRIPTION=""
      RECTANGLE= { X=90 Y=50 H=40 W=100 }
      KEY_PARAMETER=""
    }
    INITIAL_STEP="START"
    TRANSITION NAME="T1"
    {
      POSITION= { X=130 Y=100 }
      TERMINATION=F
      EXPRESSION="TRUE"
    }
    TRANSITION NAME="T2"
    {
      POSITION= { X=130 Y=220 }
      TERMINATION=F
      EXPRESSION="'OP_MELDEN:1/BSTATUS' = '$recipe_state:Complete'
AND 'OP_BEDING_DI:1/BSTATUS' = '$recipe_state:Complete'"
    }
    TRANSITION NAME="T3"
    {
      POSITION= { X=130 Y=350 }
      TERMINATION=F
      EXPRESSION="'OP_DRUCK:1/BSTATUS' = '$recipe_state:Complete'"
    }
    TRANSITION NAME="T4"
    {
      POSITION= { X=130 Y=450 }
      TERMINATION=F
      EXPRESSION="'OP_BEGAS:1/BSTATUS' = '$recipe_state:Complete'"
    }
    TRANSITION NAME="T5"
    {
      POSITION= { X=560 Y=350 }
      TERMINATION=F
      EXPRESSION="'OP_HK:1/BSTATUS' = '$recipe_state:Complete'"
    }
    TRANSITION NAME="T7"
    {
      POSITION= { X=130 Y=560 }
      TERMINATION=T
      EXPRESSION="'OP_SET_AP_B:1/BSTATUS' = '$recipe_state:Complete'
AND 'OP_SET_ALAR_15:1/BSTATUS' = '$recipe_state:Complete'
AND 'OP_KONDENS:1/BSTATUS' = '$recipe_state:Complete'"
    }
    TRANSITION NAME="T8"
    {
      POSITION= { X=340 Y=350 }
      TERMINATION=F
      EXPRESSION="'OP_SET_AKTR_25:1/BSTATUS' = '$recipe_state:Complete'"
    }
    STEP_TRANSITION_CONNECTION STEP="OP_BEDING_DI:1" TRANSITION="T2" { }
    STEP_TRANSITION_CONNECTION STEP="OP_BEGAS:1" TRANSITION="T4" { }
    STEP_TRANSITION_CONNECTION STEP="OP_DRUCK:1" TRANSITION="T3" { }
    STEP_TRANSITION_CONNECTION STEP="OP_HK:1" TRANSITION="T5" { }
    STEP_TRANSITION_CONNECTION STEP="OP_KONDENS:1" TRANSITION="T7" { }
    STEP_TRANSITION_CONNECTION STEP="OP_MELDEN:1" TRANSITION="T2" { }
    STEP_TRANSITION_CONNECTION STEP="OP_SET_AKTR_25:1" TRANSITION="T8" { }
    STEP_TRANSITION_CONNECTION STEP="OP_SET_ALAR_15:1" TRANSITION="T7" { }
    STEP_TRANSITION_CONNECTION STEP="OP_SET_AP_B:1" TRANSITION="T7" { }
    STEP_TRANSITION_CONNECTION STEP="START" TRANSITION="T1" { }
    TRANSITION_STEP_CONNECTION TRANSITION="T1" STEP="OP_MELDEN:1" { }
    TRANSITION_STEP_CONNECTION TRANSITION="T1" STEP="OP_BEDING_DI:1" { }
    TRANSITION_STEP_CONNECTION TRANSITION="T2" STEP="OP_DRUCK:1" { }
    TRANSITION_STEP_CONNECTION TRANSITION="T2" STEP="OP_SET_AKTR_25:1" { }
    TRANSITION_STEP_CONNECTION TRANSITION="T2" STEP="OP_HK:1" { }
    TRANSITION_STEP_CONNECTION TRANSITION="T3" STEP="OP_BEGAS:1" { }
    TRANSITION_STEP_CONNECTION TRANSITION="T4" STEP="OP_SET_AP_B:1" { }
    TRANSITION_STEP_CONNECTION TRANSITION="T5" STEP="OP_KONDENS:1" { }
    TRANSITION_STEP_CONNECTION TRANSITION="T8" STEP="OP_SET_ALAR_15:1" { }
  }
}`

func main() {
	fmt.Println(len(txt))
	srv := server.Server{}
	srv.Start()

}
