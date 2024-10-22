<template>
    <div class="header">
        <div class="logo">
            <img src="../assets/images/arxada-blue.png" alt="logo">
        </div>
        <div class="header-menu">
            <ul class="nav-left">
                <li><router-link to="/">Home</router-link></li>
                <li><router-link to="/operation" >Operationen</router-link></li>
                <li><router-link to="/unit">Units</router-link></li>
                <li><router-link to="/recipes">Rezepte</router-link></li>
            </ul>
            <ul class="nav-admin">
                <li><router-link to="/fhx">Daten laden</router-link></li>
                <li><router-link to="/users">Users</router-link></li>
                <li><router-link to="/plant">Betrieb</router-link></li>            
            </ul>
            <ul class="nav-right">
                <li><router-link to="/login">Login</router-link></li>
                <li><router-link to="/logout">Logout</router-link></li>
                <li><router-link to="/register">Registrieren</router-link></li>            
            </ul>
        </div>
    </div>
</template>
<script setup>
    
</script>
<style lang='css' scoped>
    .header {  
    display: flex;
    justify-content: space-between;
    align-items: center;   
    font-size: 2rem;    
    margin-right: calc(var(--padding) * 6);
    margin-left: calc(var(--padding) * 6);
    & .logo {  
        flex: 1;      
        & img {
            width: 15rem;
        }
    }
    & .header-menu {
        flex: 5;
        width: 100%;        
        display: flex;
        justify-content: space-between;
        & ul {
            display: flex;
            align-items: center;
            gap: 2rem;
            & li a {
                color: inherit;
            }
        }        
    }
}
</style>