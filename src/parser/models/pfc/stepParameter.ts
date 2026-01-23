export interface StepParameter {
    name: string;
    origin: 'DEFERRED' | 'CONSTANT' | 'INPUT' | 'OUTPUT';
    deferredTo?: string;
  }