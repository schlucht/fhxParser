import { describe, expect, it } from "vitest";
import { initalStep, parsePFC, parseStepParameter, parseSteps, parseTransition } from "../../../../src/parser/batch/pfcAlgorithm";

const txt = ['PFC_ALGORITHM',
  '{',
    'STEP NAME="DRUCK:1" DEFINITION="DRUCK"',
    '{',
      'DESCRIPTION=""',
      'RECTANGLE= { X=50 Y=130 H=40 W=100 }',
      'STEP_PARAMETER NAME="FP_BESCHREI_TEXT"',
      '{',
        'ORIGIN=DEFERRED',
        'DEFERRED_TO="FP_BESCHREI_TEXT"',
      '}',
      'STEP_PARAMETER NAME="FP_FSB_PSH"',
      '{',
        'ORIGIN=DEFERRED',
        'DEFERRED_TO="FP_FSB_PSH"',
      '}',          
      'KEY_PARAMETER=""',
    '}',
    'STEP NAME="START" DEFINITION=""',
    '{',
      'DESCRIPTION=""',
      'RECTANGLE= { X=50 Y=50 H=40 W=100 }',
      'KEY_PARAMETER=""',
    '}',
    'INITIAL_STEP="START"',
    'TRANSITION NAME="T1"',
    '{',
      'POSITION= { X=90 Y=100 }',
      'TERMINATION=F',
      'EXPRESSION="TRUE"',
    '}',
    'TRANSITION NAME="T2"',
    '{',
      'POSITION= { X=90 Y=180 }',
      'TERMINATION=T',
      `EXPRESSION='"DRUCK:1/BSTATUS' = '$phase_state:Complete'"`,
   '}',
    'STEP_TRANSITION_CONNECTION STEP="DRUCK:1" TRANSITION="T2" { }',
    'STEP_TRANSITION_CONNECTION STEP="START" TRANSITION="T1" { }',
    'TRANSITION_STEP_CONNECTION TRANSITION="T1" STEP="DRUCK:1" { }',
  '}',
'}'];

const result = {
    steps: [
        {
            name: "DRUCK:1",
            definition: "DRUCK",
            description: "",
            rectangel: { x:50, y:130, h:40, w:100, },
            keyParameter: "",
            stepParameters: [
                {
                    name: "FP_BESCHREI_TEXT",
                    origin: "DEFERRED",
                    defereTo: "FP_BESCHREI_TEXT"
                },
                {
                    name: "FP_FSB_PSH",
                    origin: "DEFERRED",
                    defereTo: "FP_FSB_PSH"
                }  
            ]

        },
        {
            name: "START",
            definition: "",
            description: "",
            rectangel: { x:50, y:50, h:40, w:100, },
            keyParameter: ""
        }
    ],
    initialStep: "START",
    transition: [
        {
            name: "T1",
            position: { x:90, y:100 },
            termination: false,
            expression: "TRUE"
        },
        {
            name: "T2",
            position: { x:90, y:180 },
            termination: true,
            expression: "'DRUCK:1/BSTATUS' = '$phase_state:Complete'",
        }
    ],
    connections: {
        stepTransition: [
            {
                step: "DRUCK:1",
                transition: "T2",
            },
            {
                step: "START",
                transition: "T1",
            }
        ],
        transitionStep: [
            {
                transition: "T1",
                step: "DRUCK:1",
            },            
        ]
    }    
}

const input2 = [
  [
    'STEP NAME="DRUCK:1" DEFINITION="DRUCK"',
    'DESCRIPTION=""',
    'RECTANGLE= { X=50 Y=130 H=40 W=100 }',
    'STEP_PARAMETER NAME="FP_BESCHREI_TEXT"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_BESCHREI_TEXT"',
    'STEP_PARAMETER NAME="FP_FSB_PSH"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_FSB_PSH"',
    'STEP_PARAMETER NAME="FP_FSB_PSL"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_FSB_PSL"',
    'STEP_PARAMETER NAME="FP_OPTION"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_OPTION"',
    'STEP_PARAMETER NAME="FP_PAH"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_PAH"',
    'STEP_PARAMETER NAME="FP_PAHH"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_PAHH"',
    'STEP_PARAMETER NAME="FP_PAL"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_PAL"',
    'STEP_PARAMETER NAME="FP_PALL"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_PALL"',
    'STEP_PARAMETER NAME="FP_PC_W"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_PC_W"',
    'STEP_PARAMETER NAME="FP_PC_W_HO"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_PC_W_HO"',
    'STEP_PARAMETER NAME="FP_PV_KENNLINIE"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_PV_KENNLINIE"',
    'STEP_PARAMETER NAME="FP_RAMPE"',
    'ORIGIN=DEFERRED',
    'DEFERRED_TO="FP_RAMPE"',
    'STEP_PARAMETER NAME="RP_FSB_WART_ZEIT"',
    'ORIGIN=CONSTANT',
    'KEY_PARAMETER=""'
  ],
  [
    'STEP NAME="START" DEFINITION=""',
    'DESCRIPTION=""',
    'RECTANGLE= { X=50 Y=50 H=40 W=100 }',
    'KEY_PARAMETER=""'
  ]
]
describe("parse.batch.pfcAlgorithm", () => {
  
  it("parse Step", () => {
    const res = parseSteps(input2);
    expect(res).toMatchObject([
      {
        name: "DRUCK:1", 
        definition: "DRUCK", 
        description: "",
        rectangle: { x:50, y:130, h:40, w:100}, 
        keyParameter: "",
      },
      {
        name: "START", 
        definition: "", 
        description: "",
        rectangle: { x:50, y:50, h:40, w:100}, 
        keyParameter: "",
      },
    ]);
  });

  it("parsed parseStepParameter", () => {
    const res = {
      name: "FP_BESCHREI_TEXT",
      origin: "DEFERRED",
      defered_to: "FP_BESCHREI_TEXT"    
    }
    const ins = [
      'STEP_PARAMETER NAME="FP_BESCHREI_TEXT"',
      'ORIGIN=DEFERRED',
      'DEFERRED_TO="FP_BESCHREI_TEXT"',
    ]
    const res2 = parseStepParameter(ins)
    expect(res2).toMatchObject({
      name: 'FP_BESCHREI_TEXT',
      origin: 'DEFERRED',
      deferredTo: 'FP_BESCHREI_TEXT'
    });
  }); 
    it("parsed initalStep", () => {
      const res = "START";
      const init = [[
        'STEP NAME="START" DEFINITION=""',
        'INITIAL_STEP="START"',
        'TRANSITION NAME="T1"'
      ]];

      const res2 = initalStep(init);
      expect(res2).toBe(res);
    });

    it("parse Transition", () => {
      const res = [
        {
          name: "T1",
          position: { x:90, y:100 },
          termination: false,
          expression: "TRUE"
        },
        {
          name: "T2",
          position: { x:90, y:180 },
          termination: true,
          expression: "'DRUCK:1/BSTATUS' = '$phase_state:Complete'",
        }
      ];
      const init = [[
        'INITIAL_STEP="START"',
        'TRANSITION NAME="T1"',        
          'POSITION= { X=90 Y=100 }',
          'TERMINATION=F',
          'EXPRESSION="TRUE"',        
        'TRANSITION NAME="T2"',        
          'POSITION= { X=90 Y=180 }',
          'TERMINATION=T',
          `EXPRESSION="'DRUCK:1/BSTATUS' = '$phase_state:Complete'"`,        
        `STEP_TRANSITION_CONNECTION STEP="DRUCK:1" TRANSITION="T2" { }`,
        `STEP_TRANSITION_CONNECTION STEP="START" TRANSITION="T1" { }`,
        `TRANSITION_STEP_CONNECTION TRANSITION="T1" STEP="DRUCK:1" { }`,
      ]];
      const res2 = parseTransition(init);
      expect(res2).toMatchObject(res);
    });
  });

  // it("Hauptresultat PFC_ALGORITHM", () => {
  //   const res = parsePFC(txt);
  //   expect(res).toMatchObject(result);
  // });
  
