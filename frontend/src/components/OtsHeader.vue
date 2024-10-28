<template>
    <div class="header w3-row">      
        <img class="header__logo w3-image w3-col l12" src="../assets/images/arxada-blue.png" alt="logo">
        
        <nav class="header__menu">
            <div class="header__menu__item">
                <router-link class="w3-button" to="/">Home</router-link>
                <router-link class="w3-button" to="/operation" >Operationen</router-link>
                <router-link class="w3-button" to="/unit">Units</router-link>
                <router-link class="w3-button" to="/recipes">Rezepte</router-link>
            </div>
            <div class="header__menu__item">
                <router-link class="w3-button" to="/fhx">Daten laden</router-link>
                <router-link class="w3-button" to="/users">Users</router-link>
                <router-link class="w3-button" to="/plant">Betrieb</router-link>            
            </div>
            <div class="header__menu__item">
                <router-link class="w3-button" to="/login">Login</router-link>
                <router-link class="w3-button" to="/logout">Logout</router-link>
                <router-link class="w3-button" to="/register">Registrieren</router-link>            
            </div>
        </nav>
    </div>
</template>
<script setup>
    
</script>
<style lang='css' scoped>

    .header {
        position: relative;
          
        display: flex;
        justify-content: space-between;
        
    }
   
   .header__logo {
        display: block;
        height: 100%;
        width: 10rem;
        
   }

   .header__menu {      
        display: flex;
        justify-content: space-between;
        align-items: center;
        
   }
 

</style>