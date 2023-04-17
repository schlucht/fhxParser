<template>
    <div class="container">
        <pre>
            <slot></slot>
        </pre>
    </div>
</template>
<script setup>

</script>
<style scoped>

    div {
        overflow-y: scroll;
        background: rgb(75, 69, 69);
        height: 80vh;
    }

    pre {     
        
        font-size: .8rem;
        
    }

</style>