import pandas as pd
import matplotlib.pyplot as plt

# Read the CSV file
df = pd.read_csv('sa.csv')

# Extract the data for plotting
generations = df['gen']
best = df['Best']
worst = df['Worst']
avg = df['Avg']

# Create the plot
plt.figure(figsize=(12, 7))

# Plot the three lines
plt.plot(generations, best, label='Best', color='green', linewidth=2, marker='o', markersize=3, markevery=max(1, len(generations)//50))
plt.plot(generations, avg, label='Average', color='blue', linewidth=2, marker='s', markersize=3, markevery=max(1, len(generations)//50))
plt.plot(generations, worst, label='Worst', color='red', linewidth=2, marker='^', markersize=3, markevery=max(1, len(generations)//50))

# Add labels and title
plt.xlabel('Generacja', fontsize=12, fontweight='bold')
plt.ylabel('Wartość Funkcji Celu', fontsize=12, fontweight='bold')
plt.title('Metaheurystyka SA: Najlepsze, średnie i najgorsze wartości funkcji w każdej iteracji', fontsize=14, fontweight='bold')

# Add grid
plt.grid(True, alpha=0.3, linestyle='--')

# Add legend
plt.legend(loc='best', fontsize=11, framealpha=0.9)

# Improve layout
plt.tight_layout()

# Save the plot
plt.savefig('sa_convergence.png', dpi=300, bbox_inches='tight')

# Show the plot
plt.show()

print("Plot saved as 'ga_convergence.png'")