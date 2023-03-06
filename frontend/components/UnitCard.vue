<template>
    <article class="unit">
        <header>            
            <a href="#">{{ ups.name }}</a>
        </header>
        <div class="body">   
            <p>{{ ups.desc }}</p>
            <p>Anzahl Parameter: <strong>{{  ups.params.length }}</strong></p>
        </div>
    </article>
</template>
<script lang="ts" setup>
import { UnitProcedure } from '~~/models/units';

    defineProps({
        ups: {
            type: Object as PropType<UnitProcedure>,
            default: null
        }
    })
</script>