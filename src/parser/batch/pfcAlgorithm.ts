import { PfcAlgorithm } from '../models/pfc/pfcAlgorithm';
import { Step } from '../models/pfc/step';
import { StepParameter } from '../models/pfc/stepParameter';
import { Transition } from '../models/pfc/transition';
import { Regex } from '../utils/const';
import { matchStrR, readBlock } from '../utils/utils';
import { parsePosition, parseRectangle } from './rectangel';

export function parsePFC(txtFhx: string[]): PfcAlgorithm | null {
	const pfcAlgo = {
		steps: [] as Step[],
		initialStep: '',
		transitions: [] as Transition[],
		connections: {
			stepTransition: [] as { step: string; transition: string }[],
			transitionStep: [] as { transition: string; step: string }[],
		},
	};
	const pfc = readBlock('STEP', txtFhx);
	// const step = readSection(pfc, "STEP", "INITIAL_STEP")
	
	pfcAlgo.steps = parseSteps(pfc) || [];
	pfcAlgo.initialStep = initalStep(pfc);
	pfcAlgo.transitions = parseTransition(pfc);
	pfcAlgo.connections = parseConnections(pfc);
	
	return pfcAlgo;
}

export function parseSteps(lines: string[][]): Step[] | null {
	const steps: Step[] = [];

	for (const l of lines) {
		const stepParams: StepParameter[] = [];
		for (let i = 0; i < l.length; i++) {
			const step = {
				name: '',
				definition: '',
				description: '',
				rectangle: { x: 0, y: 0, h: 0, w: 0 },
				keyParameter: '',
			};
			if (l[i].startsWith('STEP NAME')) {
				const n = matchStrR(Regex.step.name, l[i].trim(), 1);
			
				if (n.ok) {
					step.name = n.value;
				}
				const def = matchStrR(Regex.step.name, l[i].trim(), 2);
				if (def.ok) {
					step.definition = def.value;
				}
				const desc = matchStrR(Regex.description, l[i + 1]);
				if (desc.ok) {
					step.description = desc.value;
				}
				const rec = parseRectangle(l[i + 2]);
				if (rec) {
					step.rectangle = rec;
				}
				steps.push(step);
				continue;
			}
			if (l[i].startsWith('STEP_PARAMETER')) {
				const stepParam = parseStepParameter(l.slice(i, i + 3));
				stepParams.push(stepParam);
			}
		}
	}
	return steps;
}

export function initalStep(lines: string[][]): string {
	for (const ls of lines) {
		for (const l of ls) {
			if (l.startsWith('INITIAL_STEP')) {
				const init = matchStrR(Regex.step.initialName, l);
			
				if (init.ok) {
					return init.value;
				}
			}
		}
	}
	return '';
}

export function parseStepParameter(lines: string[]): StepParameter {
	const stepParam: StepParameter = {
		name: '',
		origin: 'CONSTANT',
		deferredTo: '',
	};
	const name = matchStrR(Regex.step.paramName, lines[0]);
	if (name.ok) {
		stepParam.name = name.value;
	}
	const origin = matchStrR(Regex.step.origin, lines[1]);
	if (origin.ok) {
		stepParam.origin = origin.value as StepParameter['origin'];
	}
	const defere = matchStrR(Regex.step.defered, lines[2]);
	if (defere.ok) {
		stepParam.deferredTo = defere.value;
	}

	return stepParam;
}

export function parseTransition(lines: string[][]): Transition[] {
	const transitions: Transition[] = [];
	for (const l of lines) {
		for(let i = 0; i < l.length; i++){
			const trans = {
				name: '',
				position: { x: 0, y: 0 },
				termination: false,
				expression: '',
			};
			if(l[i].startsWith('TRANSITION ', 0)) {
				const name = matchStrR(Regex.step.transition, l[i]);
				if(name.ok) {
					trans.name = name.value;
				}		
				
				if(l[i + 1].startsWith('POSITION')){	
					const pos = parsePosition(l[i + 1]);
					if(pos) {
						trans.position = pos;
					}
				}
				const term = matchStrR(Regex.step.termination, l[i + 2]);
				if(term.ok) {
					trans.termination = term.value === 'T' ? true : false;
				}
				const expr = matchStrR(Regex.step.expression, l[i + 3]);
				if(expr.ok) {
					trans.expression = expr.value;
				}
				transitions.push(trans);
			}
		}
	}
	return transitions;
}

export function parseConnections(lines: string[][]): {stepTransition: { step: string; transition: string }[];
transitionStep: { transition: string; step: string }[]}{
	
	const connections = {
		stepTransition: [{step: '', transition: ''}],
		transitionStep: [{transition: '', step: ''}]
	}
	
	for(const ls of lines) {
		for(const l of ls) {
			if(l.startsWith('STEP_TRANSITION_CONNECTION')) {
				const st = {step: '', transition: ''};
				const step = matchStrR(Regex.step.stepTrans, l, 1);
				if(step.ok) {
					st.step = step.value;
				}
				const trans = matchStrR(Regex.step.stepTrans, l, 2);
				if(trans.ok) {
					st.transition = trans.value;
				}
				connections.stepTransition.push(st);
			}
			if(l.startsWith('TRANSITION_STEP_CONNECTION')) {
				const st = {transition: '', step: ''};
				const trans = matchStrR(Regex.step.transStep, l, 1);
				if(trans.ok) {
					st.transition = trans.value;
				}
				const step = matchStrR(Regex.step.transStep, l, 2);
				if(step.ok) {
					st.step = step.value;
				}
				connections.transitionStep.push(st);
			}
		}
	}

	return {
		stepTransition: connections.stepTransition,
		transitionStep: connections.transitionStep
	}
}