package com.wuriyanto.dev.webinar.shared

import android.content.Context
import android.content.Context.MODE_PRIVATE
import java.util.*


private var uniqueID: String? = null
private val PREF_CLIENT_ID = "PREF_CLIENT_ID"
private val PREF_ACCESS_TOKEN = "PREF_ACCESS_TOKEN"

@Synchronized
fun getClientID(context: Context): String {
    if (uniqueID == null) {
        val sharedPrefs = context.getSharedPreferences(
            PREF_CLIENT_ID, MODE_PRIVATE
        )
        uniqueID = sharedPrefs.getString(PREF_CLIENT_ID, null)
        if (uniqueID == null) {
            uniqueID = UUID.randomUUID().toString()
            val editor = sharedPrefs.edit()
            editor.putString(PREF_CLIENT_ID, uniqueID)
            editor.commit()
        }
    }
    return uniqueID as String
}

@Synchronized
fun setAccessTokenToPreference(context: Context, accessToken: String) {
    val sharedPrefs = context.getSharedPreferences(
        PREF_ACCESS_TOKEN, MODE_PRIVATE
    )

    val editor = sharedPrefs.edit()
    editor.putString(PREF_ACCESS_TOKEN, accessToken)

    editor.commit()
}

fun removeAccessTokenToPreference(context: Context) {
    val sharedPrefs = context.getSharedPreferences(
        PREF_ACCESS_TOKEN, MODE_PRIVATE
    )

    val editor = sharedPrefs.edit()
    editor.remove(PREF_ACCESS_TOKEN)

    editor.commit()
}

fun getAccessTokenFromPreference(context: Context): String? {
    val sharedPrefs = context.getSharedPreferences(
        PREF_ACCESS_TOKEN, MODE_PRIVATE
    )
    return sharedPrefs.getString(PREF_ACCESS_TOKEN, null)
}