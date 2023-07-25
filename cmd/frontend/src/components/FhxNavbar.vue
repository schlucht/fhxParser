<template>
  <nav class="header-nav">
    <router-link to="/" alt="Startseite" class="header-nav-brand">
      <img class="header-logo" src="@/assets/img/header-logo.png" />
    </router-link>
    <router-link to="/" alt="Startseite" class="header-nav-link">
      <span class="icomoon-home5" aria-hidden="true"></span>
    </router-link>
    <router-link to="/" alt="Startseite" class="header-nav-link"> Home </router-link>
    <router-link to="/" alt="Rezepte" class="header-nav-link"> Rezepte </router-link>
    <router-link to="/" alt="Unitprocedures" class="header-nav-link"> Unitprocedures </router-link>
    <router-link to="/" alt="Startseite" class="header-nav-link"> Operations </router-link>
  </nav>
</template>
<script setup>
import { RouterLink } from 'vue-router'
</script>
<style scoped>
.header-nav {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  width: 100%;
  height: 8.5rem;
  padding-left: 0.5rem;
  background-color: var(--white);
}

.header-nav .header-nav-brand {
  height: 100%;
  margin-top: 0.5rem;
}

.header-nav .header-nav-brand .header-logo {
  height: auto;
  display: block;
  max-height: 6.5rem;
}

.header-nav .header-nav-link {
  padding: 0rem 1.5rem;
  cursor: pointer;
}

.header-nav .header-nav-link:nth-child(2) {
  margin-left: 8rem;
  color: var(--light-blue);
  font-size: 2rem;
}
</style>
