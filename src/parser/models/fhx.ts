import { Schema } from "./document/schema";
import { Locale } from "./document/locale";
import { BatchRecipe } from "./batch/batchRecipe";

// Grundbausteine
export type BoolFlag = boolean;



export interface Rect { x: number; y: number; h: number; w: number; }
export interface Pos { x: number; y: number; }

export interface Header {
  version?: string;
  dateStr?: string;
}

// Dokumentwurzel
export interface FhxDocument {
  header: Header;
  schema: Schema; 
  locale: Locale;
  recipes: BatchRecipe[];
}








