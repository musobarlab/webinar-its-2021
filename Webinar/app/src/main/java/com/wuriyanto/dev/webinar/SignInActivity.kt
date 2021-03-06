package com.wuriyanto.dev.webinar

import android.content.Intent
import android.os.Bundle
import android.util.Log
import android.view.View
import android.widget.Button
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import com.google.android.material.textfield.TextInputEditText
import com.wuriyanto.dev.webinar.model.SignIn
import com.wuriyanto.dev.webinar.model.SignInResponse
import com.wuriyanto.dev.webinar.service.SignInService
import com.wuriyanto.dev.webinar.shared.ServiceBuilder
import com.wuriyanto.dev.webinar.shared.removeAccessTokenToPreference
import com.wuriyanto.dev.webinar.shared.setAccessTokenToPreference
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

class SignInActivity : AppCompatActivity() {

    private val TAG = "SignInActivity"

    private lateinit var textEmail: TextInputEditText
    private lateinit var textPassword: TextInputEditText
    private lateinit var buttonLogin: Button
    private lateinit var buttonSignUp: Button

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_sign_in)

        textEmail = findViewById(R.id.txt_sign_in_email)
        textPassword = findViewById(R.id.txt_sign_in_password)

        buttonLogin = findViewById(R.id.btn_login)
        buttonSignUp = findViewById(R.id.btn_sign_up_navigate)

        buttonSignUp.setOnClickListener {
            startActivity(Intent(this, SignUpActivity::class.java))
            finish()
        }

        buttonLogin.setOnClickListener { login(it) }
    }

    private fun login(v: View) {

        val loginRequest = ServiceBuilder.buildService(SignInService::class.java)

        if (textEmail.text.toString().isBlank()) {
            showToast("email cannot be empty")
            return
        }

        if (textPassword.text.toString().isBlank()) {
            showToast("password cannot be empty")
            return
        }

        val signInData = SignIn(textEmail.text.toString(), textPassword.text.toString())
        Log.d(TAG, signInData.username)
        Log.d(TAG, signInData.password)

        loginRequest.signIn(signInData).enqueue(
            object : Callback<SignInResponse> {
                override fun onResponse(call: Call<SignInResponse>, response: Response<SignInResponse>) {
                    Log.d(TAG, "login succeed")
                    Log.d(TAG, response.body().toString())
                    onResult(response.body(), null)
                }

                override fun onFailure(call: Call<SignInResponse>, t: Throwable) {
                    Log.d(TAG, "login fail: ${t.message}")
                    onResult(null, t.message)
                }

            }
        )
    }

    private fun showToast(msg: String) {
        Toast.makeText(applicationContext, msg, Toast.LENGTH_SHORT).show()
    }

    private fun onResult(response: SignInResponse?, messageError: String?) {
        if (response == null) {
            if (messageError != null) {
                showToast("login failed: ${messageError}")
            } else {
                showToast("login failed")
            }
        } else {
            if (response.success) {
                // remove previous token
                removeAccessTokenToPreference(this)

                // set access token
                val accessToken = "Bearer ${response.data.accessToken}"
                setAccessTokenToPreference(this, accessToken)

                startActivity(Intent(this, ProfileActivity::class.java))
                finish()
            } else {
                showToast(response.message)
            }
        }
    }
}