<template>
    <li>        
        <slot />   
    </li>
</template>
<script setup>
    
</script>
<style lang='css' scoped>
    li {
        background-color: var(--light-gray);       
        margin-top: 10px;
        padding: 5px 15px;
        border-radius: 10px;
        box-shadow: 3px 3px 3px rgba(0,0,0,.1);
    }
</style>