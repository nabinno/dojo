module StartingOut
  ( startingOut
  ) where

startingOut :: IO ()
startingOut = do
  putStrLn (doubleNum 3)

doubleNum :: Int -> String
doubleNum x = show (x * x)
