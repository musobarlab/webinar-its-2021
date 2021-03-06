package com.wuriyanto.dev.webinar.shared


import com.google.gson.GsonBuilder

val gson = GsonBuilder().create()

inline fun <reified T> jsonToEntity(payload: String): T? {
    return gson.fromJson<T>(payload, T::class.java)
}

fun <T> entityToJson(entity: T): String? {
    return gson.toJson(entity)
}