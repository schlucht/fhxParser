<template>
    <div>
        <div class="footer"> 
            &copy;OTS - {{new Date().getMonth() + 1}} / {{new Date().getFullYear()}}               
        </div>
    </div>
</template>
<script setup>
    
</script>
<style lang='css' scoped>
    footer {
    color: var(--clr-white);
    margin-top: calc(var(--padding) * 6);
}
</style>