package com.wuriyanto.dev.webinar.model

import com.google.gson.annotations.SerializedName

data class SignUp(
    val email: String,
    val fullname: String,
    val password: String
)

data class SignUpResponse(
    @SerializedName("success") val success: Boolean,
    @SerializedName("code") val code: Int,
    @SerializedName("message") val message: String,
    @SerializedName("data") val data: SignUpResponseData
)

data class SignUpResponseData(
    val email: String,
    val fullname: String,
    val status: String,
    val rank: String,
    val bonus: Float
)