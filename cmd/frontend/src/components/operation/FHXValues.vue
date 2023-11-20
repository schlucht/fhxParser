<template>   
    <tr>
        <th>Min</th>
        <th>Value</th>
        <th>Max</th>
        <th>Unit</th>
    </tr>
    <tr>
        <td>{{ pa.values.low }}</td>
        <td>{{ pa.values.CV }}</td>
        <td>{{ pa.values.hight }}</td>
        <td>{{ pa.values.unit }}</td>
    </tr>    
</template>
<script setup>
    import { defineProps } from 'vue';
    const pa = defineProps(['values'])      
    
</script>
<style lang='css' scoped>
    td {
        background-color: rgb(198, 204, 205);
        padding: .4rem;  
        text-align: center;
        width: 25px;     
    }
    th {
        background-color: var(--light-gray);
        color: white;
    }
</style>