<template>
    <h1>Operation</h1>
    <div v-for="op of ops" :key="op.op_id">
        {{ op.opname }}
    </div>   
</template>
<script setup>
import { Operation } from '../models/operation.js';
import { store } from '../store/store.js';
import { ref } from 'vue';

const ops = ref(null);
async function loadOperation() {
    const op = new Operation(
        store.plant.value.id
    );
    await op.loadOperation();
    ops.value = op.operations;
}
loadOperation()
</script>
<style lang='css' scoped>
    
</style>