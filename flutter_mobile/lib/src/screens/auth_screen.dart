import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_mobile/src/screens/reg_screen.dart';

@RoutePage()
class AuthScreen extends StatelessWidget {
  const AuthScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 32),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              OutlinedButton(
                onPressed: () => context.router.pushNamed("/login"),
                child: Text("Login"),
              ),
              OutlinedButton(
                onPressed: () => context.router.pushNamed("/reg"),
                child: Text("Reg"),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
