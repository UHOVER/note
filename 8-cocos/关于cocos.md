####新加的c++ 绑定

```c++
例如: 新建 AssetsSingleton.cpp
新建cpp或者拖进来已经写好的cpp 
先将他放到 frameworks/runtime-src/Class, 然后在Xcode 的工程中拖进来 Added folders 选择 Create groups，在勾选 Add to targets 然后将这个文件加到 IOS 项目中

在proj.android-studio/app/jni/Android.mk 中，在LOCAL_SRC_FILES := \
../../../Classes/AssetsSingleton.cpp \
这是 Android 项目

写一个lua 绑定cpp，AssetsSingleton.cpp要绑定的lua方法必须是static的，然后写一个 lua_custom_auto.cpp 绑定类

class AssetsSingleton
{
public:
    // 热更新实例，使用单例模式是为了加载模块时减少向服务器请求的时间
    static AssetsManagerEx *getAssetsManagerEx();
};

// lua_custom_auto.cpp
int lua_AssetsSingleton_getAssetsManagerEx(lua_State* tolua_S)
{
    int argc = lua_gettop(tolua_S)-1;    
    if (argc == 0) {
        
        AssetsManagerEx* ret = AssetsSingleton::getAssetsManagerEx();
        object_to_luaval<AssetsManagerEx>(tolua_S,"cc.AssetsManagerEx",(AssetsManagerEx*)ret);
        return 1;
    }
    CCLOG("参数错误，downloadModuleByName 的参数是一个字符串");
    return 0;
}

TOLUA_API int register_custom(lua_State* tolua_S)
{
    tolua_open(tolua_S);
    tolua_module(tolua_S,"qyoo",0);
    tolua_beginmodule(tolua_S,"qyoo");
    tolua_usertype(tolua_S, "qyoo.AssetsSingleton");
    tolua_cclass(tolua_S,"AssetsSingleton", "qyoo.AssetsSingleton", "", nullptr);
    
    tolua_beginmodule(tolua_S,"AssetsSingleton");
        tolua_function(tolua_S, "getAssetsManagerEx", lua_AssetsSingleton_getAssetsManagerEx);
    
    tolua_endmodule(tolua_S);
    tolua_endmodule(tolua_S);
    std::string typeName = typeid(AssetsSingleton).name();
    g_luaType[typeName] = "qyoo.AssetsSingleton";
    g_typeCast["AssetsSingleton"] = "qyoo.AssetsSingleton";
    return 1;
}

```