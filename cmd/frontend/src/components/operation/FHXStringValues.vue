<template>   
    <tr >
        <th v-if="pa.values.value_set !== ''">Param_Set</th>        
        <th v-if="pa.values.CV == ''">Text</th>        
    </tr>
    <tr >
        <td v-if="pa.values.value_set !== ''">{{ pa.values.value_set }}</td>        
        <td v-if="pa.values.CV == ''">{{ pa.values.string_value || "&nbsp;" }}</td>        
    </tr>    
</template>
<script setup>
    import { defineProps } from 'vue';
    const pa = defineProps(['values'])
      
    
</script>
<style lang='css' scoped>
td {
        background-color: rgba(189, 226, 230,.4);
        padding: .4rem;   
        border: solid 1px rgba(189, 226, 230,.8);
        width: 50%;
        padding: 0.5rem 1.0rem;
    }
th {
    background-color: var(--light-blue);
    color: white;
}
    
</style>