<template>
  <div v-if="item.children">
    <template v-if="item.children.length == 0">
      <el-menu-item :index="item.path" @click="openMenu(item)">
        <i class="el-icon-menu"></i>
        {{ item.label }}
      </el-menu-item>
    </template>

    <el-submenu v-else :index="item.path">
      <template slot="title">
        <i class="el-icon-menu"></i>
        {{ item.label }}
      </template>

      <template v-for="child in item.children">
        <sidebar-item
          v-if="child.children && child.children.length > 0"
          :item="child"
          :key="child.path"
        />
        <el-menu-item
          v-else
          :key="child.path"
          :index="child.path"
          @click="openMenu(child)"
        >
          <i class="el-icon-location"></i>
          {{ child.label }}
        </el-menu-item>
      </template>
    </el-submenu>
  </div>
</template>

<script>
// import { mapGetters } from "vuex";
export default {
  name: "SidebarItem",
  props: {
    item: {
      type: Object,
      required: true,
    },
    activeItem: {
      type: Object,
    },
  },
  watch: {
  },
  computed: {
  },
  methods: {
    openMenu(item) {
      this.$router.push({ name: item.name });
      this.$store.commit("ADD_TAG", item);
    },
  },
};
</script>

<style>
</style>