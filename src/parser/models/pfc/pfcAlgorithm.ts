import { Step } from "./step";
import { Transition } from "./transition";

export interface PfcAlgorithm {
    steps: Step[];
    initialStep?: string;
    transitions: Transition[];
    connections: {
      stepTransition: { step: string; transition: string }[];
      transitionStep: { transition: string; step: string }[];
    };
  }
  
  