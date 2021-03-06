package com.wuriyanto.dev.webinar.shared

import okhttp3.CertificatePinner
import okhttp3.OkHttpClient
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory

private const val BASE_URL = "https://yourdomain.co/"

object ServiceBuilder {

    private val certPinner = CertificatePinner.Builder()
        .add("yourdomain.co", "sha256/lHW9w+SjBInznWN+CckwpLBAxpM0fSycqlWteqdRg1E=")
        //.add("yourdomain.co", "sha256/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=")
        .build()
    private val client = OkHttpClient.Builder()
        .certificatePinner(certPinner)
        .build()

    private val retrofit = Retrofit.Builder()
        .baseUrl(BASE_URL)
        .addConverterFactory(GsonConverterFactory.create())
        .client(client)
        .build()

    fun<T> buildService(service: Class<T>): T{
        return retrofit.create(service)
    }
}