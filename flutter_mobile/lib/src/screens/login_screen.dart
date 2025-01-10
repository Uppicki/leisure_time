import 'package:auto_route/annotations.dart';
import 'package:flutter/material.dart';
import 'package:flutter_mobile/src/features/auth/service/auth_service.dart';
import 'package:provider/provider.dart';

@RoutePage()
class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final _authService = context.read<AuthService>();

    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          onPressed: () => Navigator.pop(context),
          icon: Icon(Icons.arrow_back),
        ),
        title: Text("Login"),
      ),
      body: Center(
        child: FractionallySizedBox(
          widthFactor: 0.8,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              TextField(
                decoration: InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: "Login",
                ),
              ),
              SizedBox(
                height: 16.0,
              ),
              TextField(
                decoration: InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: "Password",
                ),
              ),
              SizedBox(
                height: 32.0,
              ),
              FilledButton(
                onPressed: () {
                  _authService.login();
                },
                child: Text("Login"),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
