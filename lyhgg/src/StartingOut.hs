module StartingOut
  ( startingOut
  ) where

startingOut :: IO ()
startingOut = do
  putStrLn ""
  putStrLn "# 2. Straight out"
  doubleNum
  putStrLn ""
  boolean

doubleNum :: IO ()
doubleNum = do
  putStrLn "## Number"
  putStrLn ("3 * 3: " <> (show (3 * 3)))

boolean :: IO ()
boolean = do
  putStrLn "## Boolean"
  put "True && False" (True && False)
  put "True && True" (True && True)
  put "False || True" (False || True)
  put "not False" (not False)
  put "not (True && True)" (not (True && True))
  put "5 == 5" (5 == 5)
  put "1 == 0" (1 == 0)
  put "5 /= 5" (5 /= 5)
  put "5 /= 4" (5 /= 4)
  put "hello == hello" ("hello" == "hello")

put :: String -> Bool -> IO ()
put x y = do
  putStrLn (x <> ": " <> (show y))
